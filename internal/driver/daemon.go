// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"

	"github.com/ory/x/otelx/semconv"

	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ory/keto/internal/namespace/namespacehandler"
	"github.com/ory/keto/internal/schema"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/logrusx"
	prometheus "github.com/ory/x/prometheusx"

	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/x/reqlog"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"

	"github.com/ory/analytics-go/v5"
	"github.com/ory/x/healthx"
	"github.com/ory/x/metricsx"
	"github.com/ory/x/otelx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/graceful"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func (r *RegistryDefault) enableSqa(cmd *cobra.Command) {
	ctx := cmd.Context()

	r.sqaService = metricsx.New(
		cmd,
		r.Logger(),
		r.Config(ctx).Source(),
		&metricsx.Options{
			Service:       "keto",
			DeploymentId:  metricsx.Hash(r.Persister().NetworkID(ctx).String()),
			IsDevelopment: strings.HasPrefix(r.Config(ctx).DSN(), "sqlite"),
			WriteKey:      "jk32cFATnj9GKbQdFL7fBB9qtKZdX9j7",
			WhitelistedPaths: []string{
				"/",
				healthx.AliveCheckPath,
				healthx.ReadyCheckPath,
				healthx.VersionPath,

				relationtuple.ReadRouteBase,
				check.RouteBase,
				expand.RouteBase,
			},
			BuildVersion: config.Version,
			BuildHash:    config.Commit,
			BuildTime:    config.Date,
			Config: &analytics.Config{
				Endpoint:             "https://sqa.ory.sh",
				GzipCompressionLevel: 6,
				BatchMaxSize:         500 * 1000,
				BatchSize:            1000,
				Interval:             time.Hour * 6,
			},
		},
	)
}

func (r *RegistryDefault) ServeAllSQA(cmd *cobra.Command) error {
	r.enableSqa(cmd)
	return r.ServeAll(cmd.Context())
}

func (r *RegistryDefault) ServeAll(ctx context.Context) error {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	serveFuncs := []func(context.Context, chan<- struct{}) func() error{
		r.serveRead,
		r.serveWrite,
		r.serveOPLSyntax,
		r.serveMetrics,
	}

	doneShutdown := make(chan struct{}, len(serveFuncs))

	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

		select {
		case <-osSignals:
			cancel()
		case <-innerCtx.Done():
		}

		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), graceful.DefaultShutdownTimeout)
		defer cancel()

		nWaitingForShutdown := len(serveFuncs)
		for {
			select {
			case <-ctx.Done():
				return
			case <-doneShutdown:
				nWaitingForShutdown--
				if nWaitingForShutdown == 0 {
					// graceful shutdown done
					return
				}
			}
		}
	}()

	eg := &errgroup.Group{}

	// We need to separate the setup (invoking the functions that return the serve functions) from running the serve
	// functions to mitigate race conditions in the HTTP router.
	for _, serve := range serveFuncs {
		eg.Go(serve(innerCtx, doneShutdown))
	}

	return eg.Wait()
}

func (r *RegistryDefault) serveRead(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.ReadRouter(ctx), r.ReadGRPCServer(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = otelx.TraceHandler(rt, otelhttp.WithTracerProvider(tracer.Provider()))
	}

	return func() error {
		addr, listenFile := r.Config(ctx).ReadAPIListenOn()
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "read"), addr, listenFile, rt, s, done)
	}
}

func (r *RegistryDefault) serveWrite(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.WriteRouter(ctx), r.WriteGRPCServer(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = otelx.TraceHandler(rt, otelhttp.WithTracerProvider(tracer.Provider()))
	}

	return func() error {
		addr, listenFile := r.Config(ctx).WriteAPIListenOn()
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "write"), addr, listenFile, rt, s, done)
	}
}

func (r *RegistryDefault) serveOPLSyntax(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.OPLSyntaxRouter(ctx), r.OplGRPCServer(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = otelx.TraceHandler(rt, otelhttp.WithTracerProvider(tracer.Provider()))
	}

	return func() error {
		addr, listenFile := r.Config(ctx).OPLSyntaxAPIListenOn()
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "opl"), addr, listenFile, rt, s, done)
	}
}

func (r *RegistryDefault) serveMetrics(ctx context.Context, done chan<- struct{}) func() error {
	ctx, cancel := context.WithCancel(ctx)

	//nolint:gosec // graceful.WithDefaults already sets a timeout
	s := graceful.WithDefaults(&http.Server{
		Handler: r.metricsRouter(ctx),
	})

	return func() error {
		defer cancel()
		eg := &errgroup.Group{}

		addr, listenFile := r.Config(ctx).MetricsListenOn()
		l, err := listenAndWriteFile(ctx, addr, listenFile)
		if err != nil {
			return err
		}

		eg.Go(func() error {
			if err := s.Serve(l); !errors.Is(err, http.ErrServerClosed) {
				return errors.WithStack(err)
			}
			return nil
		})
		eg.Go(func() (err error) {
			defer func() {
				l := r.Logger().WithField("endpoint", "metrics")
				if err != nil {
					l.WithError(err).Error("graceful shutdown failed")
				} else {
					l.Info("gracefully shutdown server")
				}
				done <- struct{}{}
			}()

			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
			defer cancel()
			return s.Shutdown(ctx)
		})

		return eg.Wait()
	}
}

func multiplexPort(ctx context.Context, log *logrusx.Logger, addr, listenFile string, router http.Handler, grpcS *grpc.Server, done chan<- struct{}) error {
	l, err := listenAndWriteFile(ctx, addr, listenFile)
	if err != nil {
		return err
	}

	m := cmux.New(l)
	m.SetReadTimeout(graceful.DefaultReadTimeout)

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1())

	//nolint:gosec // graceful.WithDefaults already sets a timeout
	restS := graceful.WithDefaults(&http.Server{
		Handler: router,
	})

	eg := &errgroup.Group{}

	eg.Go(func() error {
		if err := grpcS.Serve(grpcL); !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := restS.Serve(httpL); !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		err := m.Serve()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			// unexpected error
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() (err error) {
		defer func() {
			if err != nil {
				log.WithError(err).Error("graceful shutdown failed")
			} else {
				log.Info("gracefully shutdown server")
			}
			done <- struct{}{}
		}()

		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), graceful.DefaultShutdownTimeout)
		defer cancel()

		shutdownEg := errgroup.Group{}
		shutdownEg.Go(func() error {
			// we ignore net.ErrClosed, because a cmux listener's close func is actually the one of the root listener (which is closed in a racy fashion)
			if err := restS.Shutdown(ctx); !(err == nil || errors.Is(err, http.ErrServerClosed) || errors.Is(err, net.ErrClosed)) {
				// unexpected error
				return errors.WithStack(err)
			}
			return nil
		})
		shutdownEg.Go(func() error {
			gracefulDone := make(chan struct{})
			go func() {
				grpcS.GracefulStop()
				close(gracefulDone)
			}()
			select {
			case <-gracefulDone:
				return nil
			case <-ctx.Done():
				grpcS.Stop()
				return errors.New("graceful stop of gRPC server timed out, had to force it")
			}
		})

		return shutdownEg.Wait()
	})

	return eg.Wait()
}

func (r *RegistryDefault) allHandlers() []Handler {
	if len(r.handlers) == 0 {
		r.handlers = []Handler{
			relationtuple.NewHandler(r),
			check.NewHandler(r),
			expand.NewHandler(r),
			namespacehandler.New(r),
			schema.NewHandler(r),
		}
	}
	return r.handlers
}

func listenAndWriteFile(ctx context.Context, addr, listenFile string) (net.Listener, error) {
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", addr)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("unable to listen on %q: %w", addr, err))
	}
	const filePrefix = "file://"
	if strings.HasPrefix(listenFile, filePrefix) {
		if err := os.WriteFile(listenFile[len(filePrefix):], []byte(l.Addr().String()), 0600); err != nil {
			return nil, errors.WithStack(fmt.Errorf("unable to write listen file %q: %w", listenFile, err))
		}
	}
	return l, nil
}

func (r *RegistryDefault) ReadRouter(ctx context.Context) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "read#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	br := &x.ReadRouter{Router: httprouter.New()}
	r.PrometheusManager().RegisterRouter(br.Router)
	r.MetricsHandler().SetRoutes(br.Router)

	r.HealthHandler().SetHealthRoutes(br.Router, false)
	r.HealthHandler().SetVersionRoutes(br.Router)

	for _, h := range r.allHandlers() {
		if h, ok := h.(ReadHandler); ok {
			h.RegisterReadRoutes(br)
		}
	}

	n.Use(r.PrometheusManager())
	n.UseHandler(br)

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("read")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) WriteRouter(ctx context.Context) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	pr := &x.WriteRouter{Router: httprouter.New()}
	r.PrometheusManager().RegisterRouter(pr.Router)
	r.MetricsHandler().SetRoutes(pr.Router)

	r.HealthHandler().SetHealthRoutes(pr.Router, false)
	r.HealthHandler().SetVersionRoutes(pr.Router)

	for _, h := range r.allHandlers() {
		if h, ok := h.(WriteHandler); ok {
			h.RegisterWriteRoutes(pr)
		}
	}

	n.Use(r.PrometheusManager())
	n.UseHandler(pr)

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("write")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) OPLSyntaxRouter(ctx context.Context) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "syntax#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	pr := &x.OPLSyntaxRouter{Router: httprouter.New()}
	r.PrometheusManager().RegisterRouter(pr.Router)
	r.MetricsHandler().SetRoutes(pr.Router)

	r.HealthHandler().SetHealthRoutes(pr.Router, false)
	r.HealthHandler().SetVersionRoutes(pr.Router)

	for _, h := range r.allHandlers() {
		if h, ok := h.(OPLSyntaxHandler); ok {
			h.RegisterSyntaxRoutes(pr)
		}
	}

	n.Use(r.PrometheusManager())
	n.UseHandler(pr)

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("write")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) grpcRecoveryHandler(p interface{}) error {
	r.Logger().
		WithField("reason", p).
		WithField("stack_trace", string(debug.Stack())).
		WithField("handler", "rate_limit").
		Error("panic recovered")
	return status.Errorf(codes.Internal, "%v", p)
}

func (r *RegistryDefault) unaryInterceptors(ctx context.Context) []grpc.UnaryServerInterceptor {
	is := []grpc.UnaryServerInterceptor{
		grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandler(r.grpcRecoveryHandler)),
	}
	is = append(is, r.defaultUnaryInterceptors...)
	is = append(is,
		herodot.UnaryErrorUnwrapInterceptor,
		grpcLogrus.UnaryServerInterceptor(InterceptorLogger(r.l.Logrus())),
		r.pmm.UnaryServerInterceptor,
	)
	if r.sqaService != nil {
		is = append(is, r.sqaService.UnaryInterceptor)
	}
	return is
}

func (r *RegistryDefault) streamInterceptors(ctx context.Context) []grpc.StreamServerInterceptor {
	is := []grpc.StreamServerInterceptor{
		grpcRecovery.StreamServerInterceptor(grpcRecovery.WithRecoveryHandler(r.grpcRecoveryHandler)),
	}
	is = append(is, r.defaultStreamInterceptors...)
	is = append(is,
		herodot.StreamErrorUnwrapInterceptor,
		grpcLogrus.StreamServerInterceptor(InterceptorLogger(r.l.Logrus())),
		r.pmm.StreamServerInterceptor,
	)
	if r.sqaService != nil {
		is = append(is, r.sqaService.StreamInterceptor)
	}
	return is
}

func (r *RegistryDefault) newGrpcServer(ctx context.Context) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.ChainStreamInterceptor(r.streamInterceptors(ctx)...),
		grpc.ChainUnaryInterceptor(r.unaryInterceptors(ctx)...),
	}
	if r.Tracer(ctx).IsLoaded() {
		opts = append(opts,
			grpc.StatsHandler(grpcOtel.NewServerHandler(grpcOtel.WithTracerProvider(otel.GetTracerProvider()))),
		)
	}

	opts = append(opts, r.defaultGRPCServerOptions...)
	if r.grpcTransportCredentials != nil {
		opts = append(opts, grpc.Creds(r.grpcTransportCredentials))
	}
	server := grpc.NewServer(opts...)
	return server
}

func (r *RegistryDefault) ReadGRPCServer(ctx context.Context) *grpc.Server {
	s := r.newGrpcServer(ctx)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	rts.RegisterVersionServiceServer(s, r)
	reflection.Register(s)
	for _, h := range r.allHandlers() {
		if h, ok := h.(ReadHandler); ok {
			h.RegisterReadGRPC(s)
		}
	}
	r.pmm.Register(s)

	return s
}

func (r *RegistryDefault) WriteGRPCServer(ctx context.Context) *grpc.Server {
	s := r.newGrpcServer(ctx)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	rts.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		if h, ok := h.(WriteHandler); ok {
			h.RegisterWriteGRPC(s)
		}
	}
	r.pmm.Register(s)

	return s
}

func (r *RegistryDefault) OplGRPCServer(ctx context.Context) *grpc.Server {
	s := r.newGrpcServer(ctx)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	rts.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		if h, ok := h.(OPLSyntaxHandler); ok {
			h.RegisterSyntaxGRPC(s)
		}
	}
	r.pmm.Register(s)

	return s
}

func (r *RegistryDefault) metricsRouter(ctx context.Context) http.Handler {
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.Logger(), "keto").ExcludePaths(prometheus.MetricsPrometheusPath))
	router := httprouter.New()

	r.PrometheusManager().RegisterRouter(router)
	r.MetricsHandler().SetRoutes(router)
	n.Use(r.PrometheusManager())
	n.UseHandler(router)

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("metrics")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}
	return handler
}

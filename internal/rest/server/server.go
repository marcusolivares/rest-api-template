package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"rest-api-template/internal/config"
	openapihandler "rest-api-template/internal/openapi"
	calchandler "rest-api-template/internal/rest/handler/calculator"
	healthhandler "rest-api-template/internal/rest/handler/healthcheck"
	"rest-api-template/internal/rest/middleware"
	pgstore "rest-api-template/internal/store/pg"
	"rest-api-template/resources/gen"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// Serve wires the chi router, registers all routes and middlewares, then starts
// the HTTP server.  It blocks until a SIGINT or SIGTERM is received and then
// performs a graceful shutdown.
func Serve(cfg *config.Env, logger zerolog.Logger, pool *pgxpool.Pool) error {
	// -----------------------------------------------------------------
	// Dependencies
	// -----------------------------------------------------------------
	txStore := pgstore.NewCalculatorTxStore(pool)

	// -----------------------------------------------------------------
	// Router
	// -----------------------------------------------------------------
	router := chi.NewRouter()
	router.Use(
		chimw.RequestID,                  // injects X-Request-Id header
		middleware.ZerologLogger(logger), // structured request logging
		middleware.PanicRecovery(logger), // catches panics → JSON 500
	)

	// Health probe — no api-key required
	router.Get("/healthcheck", healthhandler.Healthcheck())

	// Calculator routes — api-key middleware is scoped to this group only
	router.Route("/calculator/v1", func(r chi.Router) {
		r.Use(middleware.APIKey(cfg.APIKey))
		r.Post("/add", calchandler.Add(logger, txStore))
		r.Post("/multiply", calchandler.Multiply(logger, txStore))
	})

	// OpenAPI spec + Redoc UI (opt-in via OPENAPI_UI_ENABLE=true)
	if cfg.OpenAPIUIEnable {
		router.Get("/openapi.json", openapihandler.SpecHandler(gen.OpenAPISpec, "openapi/openapi.json"))
		router.Get(cfg.OpenAPIUIPath, openapihandler.UIHandler())
		logger.Info().Msgf("redoc ui available at %s%s", cfg.ServerAddr(), cfg.OpenAPIUIPath)
	}

	// -----------------------------------------------------------------
	// HTTP server
	// -----------------------------------------------------------------
	srv := &http.Server{
		Addr:         cfg.ServerAddr(),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout(),
		WriteTimeout: cfg.WriteTimeout(),
	}

	serverErr := make(chan error, 1)
	go func() {
		logger.Info().Msgf("http server listening on %s", cfg.ServerAddr())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
		close(serverErr)
	}()

	// -----------------------------------------------------------------
	// Graceful shutdown
	// -----------------------------------------------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return err
	case sig := <-quit:
		logger.Info().Msgf("signal %s received, shutting down gracefully", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout())
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	logger.Info().Msg("http server stopped")
	return nil
}

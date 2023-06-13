package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/Poomipat-Ch/kenta-bot/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunHTTPServer(logger *zerolog.Logger, createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), logger, createHandler)
}

func RunHTTPServerOnAddr(addr string, logger *zerolog.Logger, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter, logger)

	rootRouter := chi.NewRouter()

	// mouting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	log.Info().Msg("Starting HTTP server on " + addr)

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to start HTTP server")
	}
}

func setMiddlewares(router *chi.Mux, logger *zerolog.Logger) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logger))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)
	// addAuthMiddleare(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	allowdOrigin := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowdOrigin) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowdOrigin,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}

// TODO: Auth middleware
// func AddAuthMiddleware(router *chi.Mux){
// 	if mockAuth, _ : strconv.ParseBool(os.Getenv("MOCK_AUTH")); mockAuth {
// 		router.Use(auth.HttpMockMiddleware)
// 		return
// 	}

// 	var opts []option.ClientOption
// }

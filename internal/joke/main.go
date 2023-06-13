package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Poomipat-Ch/kenta-bot/internal/common/logs"
	"github.com/Poomipat-Ch/kenta-bot/internal/common/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	log := logs.Init()

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		// go loadFixtures()

		server.RunHTTPServer(log, func(router chi.Router) http.Handler {
			return HandlerFromMux(HttpServer{}, router)
		})
	case "grpc":
		// TODO: implement
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}

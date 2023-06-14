package main

import (
	"context"
	"net/http"

	"github.com/Poomipat-Ch/kenta-bot/internal/common/logs"
	"github.com/Poomipat-Ch/kenta-bot/internal/common/server"
	"github.com/Poomipat-Ch/kenta-bot/internal/users/app"
	"github.com/Poomipat-Ch/kenta-bot/internal/users/ports"
	"github.com/go-chi/chi/v5"
)

func main() {
	log := logs.Init()

	ctx := context.Background()
	app := app.NewApplication(ctx)

	server.RunHTTPServer(log, func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}

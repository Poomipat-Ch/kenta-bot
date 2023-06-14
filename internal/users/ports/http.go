package ports

import (
	"net/http"

	"github.com/Poomipat-Ch/kenta-bot/internal/users/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{
		app: app,
	}
}

func (h HttpServer) GetMe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is me"))
}

func (h HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is login"))
}

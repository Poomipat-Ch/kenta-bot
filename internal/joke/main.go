package main

import (
	chai "github.com/go-chai/chai/chi"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("api/v1", func(r chi.Router) {

	})
}

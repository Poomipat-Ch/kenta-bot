package main

import (
	"net/http"

	"github.com/go-chi/render"
)

type HttpServer struct {
}

func (h HttpServer) GetJoke(w http.ResponseWriter, r *http.Request) {
	// authUser, err := auth.UserFromCtx(r.Context())
	// if err != nil {
	// 	httperr.RespondWithSlugError(err, w, r)
	// 	return
	// }

	// _, _, err := net.SplitHostPort(r.RemoteAddr)
	// if err == nil {
	// err = h.db.UpdateLastIP(r.Context(), authUser.UUID, host)
	// if err != nil {
	// 	httperr.InternalError("internal-server-error", err, w, r)
	// 	return
	// }
	// }

	// user, err := h.db.GetUser(r.Context(), authUser.UUID)
	// if err != nil {
	// 	httperr.InternalError("cannot-get-user", err, w, r)
	// 	return
	// }

	// userResponse := User{
	// 	DisplayName: authUser.DisplayName,
	// 	Balance:     user.Balance,
	// 	Role:        authUser.Role,
	// }
	jokeResponse := Joke{
		Message: "Hello World",
	}

	render.Respond(w, r, jokeResponse)
}

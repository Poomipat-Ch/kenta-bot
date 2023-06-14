// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package users

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// User defines model for User.
type User struct {
	Displayname *string `json:"displayname,omitempty"`
	Email       string  `json:"email"`
	Picture     *string `json:"picture,omitempty"`
	Username    string  `json:"username"`
}

// UserToken defines model for UserToken.
type UserToken struct {
	ExpriesIn int64  `json:"expries_in"`
	Token     string `json:"token"`
}

// LoginJSONBody defines parameters for Login.
type LoginJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody

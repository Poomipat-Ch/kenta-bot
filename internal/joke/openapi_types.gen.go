// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package main

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Joke defines model for Joke.
type Joke struct {
	Message string `json:"message"`
}

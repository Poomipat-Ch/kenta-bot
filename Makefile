include .env
export

.PHONY: openapi
openapi: openapi_http


.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh joke internal/joke main
	@./scripts/openapi-http.sh users internal/users/ports ports

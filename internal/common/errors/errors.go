package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnkown        = ErrorType{"unknown"}
	ErrorTypeAuthorization = ErrorType{"authorization"}
	ErrorTypeInvalidInput  = ErrorType{"invalid_input"}
)

type SlugError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) Type() ErrorType {
	return s.errorType
}

func NewSlugError(error string, slug string, errorType ErrorType) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: errorType,
	}
}

func NewAuthorizationError(error string, slug string) SlugError {
	return NewSlugError(error, slug, ErrorTypeAuthorization)
}

func NewInvalidInputError(error string, slug string) SlugError {
	return NewSlugError(error, slug, ErrorTypeInvalidInput)
}

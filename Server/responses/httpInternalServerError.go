package responses

// HTTPInternalServerError is for 500 OK responses
type HTTPInternalServerError struct {
	StatusCode   int
	ErrorMessage string
}

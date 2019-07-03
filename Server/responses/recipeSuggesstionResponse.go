package responses

// RecipeSuggesstionResponse is for 200 OK responses
type RecipeSuggesstionResponse struct {
	StatusCode     int
	SuccessMessage string
	Item           string
}

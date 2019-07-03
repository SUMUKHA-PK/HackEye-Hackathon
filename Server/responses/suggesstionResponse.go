package responses

import (
	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/util"
)

// SuggesstionsResponse is for 200 OK responses
type SuggesstionsResponse struct {
	StatusCode     int
	SuccessMessage string
	Data           []util.Recipe
}

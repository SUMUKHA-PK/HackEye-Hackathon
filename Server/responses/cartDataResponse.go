package responses

import (
	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/util"
)

// CartDataResponse is for successful cart response responses
type CartDataResponse struct {
	StatusCode     int
	SuccessMessage string
	CartData       util.MLPostRes
}

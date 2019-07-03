package routing

import "github.com/SUMUKHA-PK/HackEye-Hackathon/Server/util"

func GetItemFromResponse(List []util.Prediction) string {
	var (
		max = -0.1
		j   = 0
	)
	for i := range List {
		if List[i].Probability > max {
			max = List[i].Probability
			j = i
		}
	}
	return List[j].TagName
}

package routing

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GoogleAssitantWebHook is the function that receives POST req from DialogFlow
func GoogleAssistantWebHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))
}

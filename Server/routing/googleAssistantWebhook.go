package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/database"
	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/responses"
	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/util"
)

// GoogleAssistantWebHook is the function that receives POST req from DialogFlow
func GoogleAssistantWebHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))

	items, err := getData()
	if err != nil {
		log.Printf("Coudn't get data from the DialogFlow JSON in routing/googleAssistantWebhook.go/GoogleAssistantWebHook")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newReq util.AddToCartReq
	userID := "SCAMS"
	itemIDs := []string{"1a", "2c", "2b"}
	newReq.ItemIDList = itemIDs
	newReq.UserID = userID
	newReq.ItemList = items
	//add the data to the DB
	err = database.AddGroceryListToDatabase(newReq)
	if err != nil {
		log.Printf("Can't add to DB in routing/groceries.go/AddItemsToCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outData := &responses.HTTPStatusOK{200, "Successfully added data to DB"}
	outJSON, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/groceries.go/AddItemsToCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}

func getData() ([]string, error) {
	arr := []string{"apples", "bananas", "mangoes"}
	return arr, nil
}

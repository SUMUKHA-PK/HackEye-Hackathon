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

// AddItemsToCart adds items to cart
func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	// get data from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/AddItemsToCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtain the JSON data in a struct
	var newReq util.AddToCartReq
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/groceries.go/AddItemsToCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(newReq)

	//add the data to the DB
	err := database.AddGroceryListToDatabase(newReq)
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

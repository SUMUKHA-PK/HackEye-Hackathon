package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
	if len(newReq.ItemIDList) != len(newReq.ItemList) {
		http.Error(w, "Size of item list and item ID list do not match!", http.StatusInternalServerError)
		return
	}

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

// CheckOutAtHome handles the checkouts that happen at home
func CheckOutAtHome(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8008/addItemsToCart"

	payload := strings.NewReader("{\"UserID\":\"sumukha\",\"ItemList\":[\"apple\",\"banana\",\"mango\"],\"ItemIDList\":[\"1a\",\"1b\",\"3d\"]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "50e95c17-7078-4bb8-8c53-21ccaca1e04a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// CheckOutAtStore handles the checkouts that happen at home
func CheckOutAtStore(w http.ResponseWriter, r *http.Request) {

}

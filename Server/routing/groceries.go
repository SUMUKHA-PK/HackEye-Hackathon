package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

// GetItemsFromCart retrieves items from the DB
func GetItemsFromCart(w http.ResponseWriter, r *http.Request) {
	// get data from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtain the JSON data in a struct
	var newReq util.GetCartReq
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(newReq)

	data, err := database.GetDataFromDB(newReq.UserID)
	if err != nil {
		log.Printf("Coudn't get cart data from DB in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Added the cart thats already added
	var cartData util.MLPostRes
	for i := range data {
		var temp util.CartItem
		temp.Item = data[i].Item
		temp.ItemID = data[i].ItemID
		cartData.AddedCartList = append(cartData.AddedCartList, temp)
	}

	// add recommendations
	var newReq2 util.MLGetReq
	err = json.Unmarshal([]byte(RawData), &newReq2)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/groceries.go/GetItemsForCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cartData.Recipes = newReq2.Recipes

	mlCuratedlist := rankData(cartData.AddedCartList, cartData.Recipes)

	// add the predictions
	cartData.PredictedCartList = mlCuratedlist[0].MissingItems

	cartData.Recipes = mlCuratedlist

	outData := &responses.CartDataResponse{200, "Successfully obtained cart data and predictions", cartData}
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
	// The static IP of the ML server
	// url := "http://localhost:52525/pasteData"

	// payload := strings.NewReader("{ \"ExpirationTime\":\"20\",\"PasteContent\":\"e eno sssassa \", \"CustomURL\": \"soonarmoonar\"}")

	// req, _ := http.NewRequest("POST", url, payload)

	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("cache-control", "no-cache")
	// // req.Header.Add("Postman-Token", "50e95c17-7078-4bb8-8c53-21ccaca1e04a")

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Printf("Error in recieving data from ML server in routing/groceries.go/AddItemsToCart")
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))

	// get data from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/CheckOutAtHome")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtain the JSON data in a struct
	var newReq util.MLGetReq
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/groceries.go/CheckOutAtHome")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outData := &responses.SuggesstionsResponse{200, "Successfully received suggesstions from the recommendation server", newReq.Recipes}
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

// CheckOutAtStore handles the checkouts that happen at home
func CheckOutAtStore(w http.ResponseWriter, r *http.Request) {

}

var RawData string = `
{
	"Recipes" : [
		{
			"recipe_name": "Chocolate Milkshake",
			"items": ["cocoa","milk","sugar"],
			"junkness" : 1,
			"preferance" : 0,
			"recipe_number": 1
		},
		{
			"recipe_name": "Cheese Sticks",
			"items": ["salt","wheat flour","butter","cheese"],
			"junkness" : 2,
			"preferance" : 0,
			"recipe_number": 2
		},
		{
			"recipe_name": "Potato Crisps",
			"items": ["vegetable oil","potatoes","wheat flour","salt"],
			"junkness" : 2,
			"preferance" : 1,
			"recipe_number": 3
		},
		{
			"recipe_name": "Peanut Cluster",
			"items": ["peanuts","sugar","milk","soy lecithin"],
			"junkness" : 0,
			"preferance" : 0,
			"recipe_number": 4
		},
		{
			"recipe_name": "Mixed Vegetables",
			"items": ["onions","potatoes","carrots","salt","vegetable oil"],
			"junkness" : 0,
			"preferance" : 1,
			"recipe_number": 5
		},
		{
			"recipe_name": "Fried Squid",
			"items": ["squid","salt","vegetable oil"],
			"junkness" : 1,
			"preferance" : 0,
			"recipe_number": 6
		},
		{
			"recipe_name": "Grilled Sandwich",
			"items": ["bread","butter","cheese","cucumber","tomato"],
			"junkness" : 1,
			"preferance" : 0,
			"recipe_number": 7
		},
		{
			"recipe_name": "Chicken Biriyani",
			"items": ["rice","salt","vegetable oil","chicken","pepper"],
			"junkness" : 2,
			"preferance" : 0,
			"recipe_number": 8
		},
		{
			"recipe_name": "Tomato Soup",
			"items": ["tomatoes","salt","pepper","water","basil"],
			"junkness" : 0,
			"preferance" : 1,
			"recipe_number": 9
		},
		{
			"recipe_name": "Potato Wedges",
			"items": ["potatoes","salt","pepper","vegetable oil"],
			"junkness" : 2,
			"preferance" : 1,
			"recipe_number": 10
		}
	
	
	]
}
`

// GetImageResponse recognise data from the image
func GetImageResponse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/CheckOutAtHome")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type urlReq struct {
		UserID string
		URL    string
	}

	var newReq1 urlReq
	err = json.Unmarshal(body, &newReq1)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/groceries.go/CheckOutAtHome")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// The image recognition server
	url := "https://southcentralus.api.cognitive.microsoft.com/customvision/v3.0/Prediction/accf9d3c-0d15-4967-97fa-428640c2cf37/classify/iterations/wmgr1/url"

	payload := strings.NewReader("{\"Url\": \"" + newReq1.URL + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Prediction-Key", "5b6a7662b4b24116a58d16683e0606b0")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq util.ImageRecognitionRes
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/groceries.go/CheckOutAtHome")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	foodItem := GetItemFromResponse(newReq.Predictions)

	if foodItem == "sphagetti" {
		foodItem = "spaghetti"
	}

	var hashmap = make(map[string][]string)

	hashmap["pizza"] = []string{"wheat flour", "yeast", "tomato puree", "cheese", "oregano"}
	hashmap["nachos"] = []string{"Tortilla chip", "cheese", "salsa", "pinto bean", "jalapeno"}
	hashmap["spaghetti"] = []string{"ground grain", "water", "tomato sauce", "cheese"}
	hashmap["tacos"] = []string{"canola oil", "onion", "red pepper", "cheese", "taco shell"}
	hashmap["pancakes"] = []string{"batter", "maple syrup", "baking powder", "butter"}
	//get items for the item picture

	possibleMissingItems := hashmap[strings.ToLower(foodItem)]

	user := newReq1.UserID
	cartItems, err := database.GetDataFromDB(user)
	if err != nil {
		log.Printf("Coudn't get cart data from DB in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	missingItems := getmissingitems(cartItems, possibleMissingItems)
	var itemIDlist []string
	for i := range missingItems {
		itemIDlist = append(itemIDlist, "ID"+strconv.Itoa(i))
		fmt.Print(i)

	}
	fmt.Println(missingItems)
	err = database.AddGroceryListToDatabase(util.AddToCartReq{user, missingItems, itemIDlist})

	outData := &responses.RecipeSuggesstionResponse{200, "Missing items added to cart", foodItem, missingItems}
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

func rankData(AddedList []util.CartItem, List []util.Recipe) []util.Recipe {
	var result []util.Recipe
	// var (
	// 	c     = 0
	// 	index = 0
	// 	max   = -1
	// )
	// fmt.Println(AddedList)
	// fmt.Println(List)
	// for i := range List {
	// 	templist := List[i].MissingItems
	// 	c = 0
	// 	for j := range templist {
	// 		for k := range AddedList {
	// 			if templist[j].Item == AddedList[k].Item {
	// 				c += 1
	// 			}
	// 		}
	// 	}
	// 	if c > max {
	// 		c = max
	// 		index = i
	// 	}
	// }
	// result = List[index].MissingItems
	// result = append(result, util.Recipe{"newRecipe", List[index].MissingItems})
	result = append(result, util.Recipe{"a1", []util.CartItem{util.CartItem{"moo1", "mooo"}, util.CartItem{"moo1", "mooo"}}})
	result = append(result, util.Recipe{"a1", []util.CartItem{util.CartItem{"moo2", "mooo"}, util.CartItem{"moo1", "mooo"}}})
	result = append(result, util.Recipe{"a1", []util.CartItem{util.CartItem{"moo3", "mooo"}, util.CartItem{"moo1", "mooo"}}})
	return result
}

func getmissingitems(Cart []util.CartData, Missing []string) []string {
	var cartlist []string
	for i := range Cart {
		cartlist = append(cartlist, Cart[i].Item)
	}
	var returnList []string
	for i := range Missing {
		jk := 0
		for j := range cartlist {
			if Missing[i] == cartlist[j] {
				jk = 1
			}
		}
		if jk == 0 {
			returnList = append(returnList, Missing[i])
		}
	}
	return returnList
}

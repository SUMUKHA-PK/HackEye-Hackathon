package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes
func SetupRouting(r mux.Router) mux.Router {
	r.HandleFunc("/", HomeRouter).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/addItemsToCart", AddItemsToCart).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/getItemsFromCart", GetItemsFromCart).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/home/checkout", CheckOutAtHome).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/store/checkout", CheckOutAtStore).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/gaWebhook", GoogleAssistantWebHook).Methods(http.MethodPost, http.MethodOptions)
	return r
}

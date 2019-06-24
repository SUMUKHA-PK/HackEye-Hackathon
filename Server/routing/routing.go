package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes
func SetupRouting(r mux.Router) mux.Router {
	r.HandleFunc("/", HomeRouter).Methods(http.MethodGet, http.MethodOptions)
	// r.HandleFunc("/pasteData", PasteData).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/p/{id:[0-9a-zA-Z]+}", ReadPaste)
	return r
}

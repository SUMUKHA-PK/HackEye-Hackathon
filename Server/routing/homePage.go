package routing

import (
	"net/http"
)

// HomeRouter redirects to instruction page of how to work with this
func HomeRouter(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/SUMUKHA-PK/HackEye-Hackathon/blob/master/README.md", 301)
}

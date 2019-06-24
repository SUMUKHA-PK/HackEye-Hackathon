package main

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	*r = routing.SetupRouting(*r)
	server.Server(r, "8008")
}

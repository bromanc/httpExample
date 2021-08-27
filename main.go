package main

import (
	"httpExample/api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

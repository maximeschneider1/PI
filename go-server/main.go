package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/LogToLeboncoin", handler.HandleGoogleLogin)

	log.Fatal(http.ListenAndServe(":8081", nil))
}


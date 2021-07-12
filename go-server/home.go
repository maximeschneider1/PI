package handler

import (
	"fmt"
	"net/http"
)
const htmlIndex = `<html><body><a href="/GoogleLogin">Log to Leboncoin</a></body></html>`

//HandleMain display home html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}
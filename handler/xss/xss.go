package main

import (
	"fmt"
	"io"
	"net/http"
)

type Handler struct{
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	output := "<h1>Hello, " + input + "</h1>"
	io.WriteString(w, output)
}

func AvoidXSSAttack() {
	handler := &Handler{}

	// Setup the server
	http.HandleFunc("/", handler.AddHandler)
	
	// Start the server
	fmt.Println("listening on port 8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func main() {
	AvoidXSSAttack()
}
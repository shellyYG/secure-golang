package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Handler struct{
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	// Below will create a template that will be used to render the output
	// This is safe because the template is not executed with user input
	tmp, err := template.New("output").Parse("<h1>Hello :D, {{ . }}</h1>")
	if err != nil {
		fmt.Printf("Error parsing template: %s\n", err)
		return
	}
	// Below in tmp.Execute, we are executing the template with user input, but as we can see, 
	// the template is not using the user input directly, it is just using the user input as a variable
	// so the user input is not executed as code
	err = tmp.Execute(w, input)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
		return
	}
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
package check_is_number

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"secure-golang/pkg/number"
)

type Handler struct{
	number *number.Number
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "No http request param matching 'input'", http.StatusBadRequest)
		return
	}

	isNumber, _ := h.number.IsNumber(input)

	payload := make(map[string]bool)
	payload["isNumber"] = isNumber

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("An error occured while marshalling the response. Err: %s", err)
	}

	w.Write(body)
}


func Is_Number() {
	handler := &Handler{
		number: number.NewNumber(),
	}

	http.HandleFunc("/isnumber", handler.AddHandler)
	fmt.Println("listening on port 8082")
}
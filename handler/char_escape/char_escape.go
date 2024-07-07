package char_escape

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"secure-golang/pkg/char"
)

type Handler struct{
	char *char.Char
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "No http request param matching 'text'", http.StatusBadRequest)
		return
	}

	result := h.char.Escape(text)
	payload := make(map[string]string)
	payload["escapedText"] = result

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("An error occured while marshalling the response. Err: %s", err)
	}

	w.Write(body)
}


func Char_escape() {
	handler := &Handler{
		char: char.Newchar(),
	}

	http.HandleFunc("/escape", handler.AddHandler)
	fmt.Println("listening on port 8082")
}
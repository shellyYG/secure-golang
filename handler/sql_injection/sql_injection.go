package main

import (
	"database/sql"
	"net/http"
)

type Patient struct {
	Name string
	Surname string
	Age int
	Gender string
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/globomantics
	")
	db.Prepare("SELECT * FROM patients WHERE id = ?")
	
}
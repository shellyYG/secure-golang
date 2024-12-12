package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("...welcome to run tls.go...")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("This is an example server.\n"))
	})
	// yourCert.pem - path to your server certificate in PEM format
	// yourKey.pem - path to your server private key in PEM format
	log.Fatal(http.ListenAndServeTLS(":443", "yourCert.pem", "yourKey.pem", nil))
}
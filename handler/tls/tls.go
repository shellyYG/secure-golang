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
		
		// The below line force client browser to upgrade all the requests from HTTP to HTTPS, if it's not already HTTPS
		// Server adds the Strict-Transport-Security (HSTS) header to the HTTP response.
		// Sets the HSTS policy to last for 63072000 seconds (2 years).
		// The browser stores this policy and applies it to all future requests to the domain (and subdomains, because of includeSubDomains).
		// If the user attempts to access the site or its subdomains using http://, the browser automatically upgrades the request to https:// before sending it.
		// If the subdomain or main domain does not support HTTPS, the browser will block the request and show an error to the user (e.g., "This site can't be reached").
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	})
	// yourCert.pem - path to your server certificate in PEM format
	// yourKey.pem - path to your server private key in PEM format
	log.Fatal(http.ListenAndServeTLS(":443", "yourCert.pem", "yourKey.pem", nil))
}
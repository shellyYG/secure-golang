// SNI: Server Name Indication
// This is DEMO only, it won't work since I don't have certificate path

package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

type Certificate struct{
	CertFile string
	Keyfile string
}

func main() {
	fmt.Println("Demo SNI....")
	httpsServer := &http.Server{
		Addr: ":8080",
	}
	var certs []Certificate
	certs = append(certs, Certificate{
		CertFile: "../etc/yourSite.pem", // Your site certificate key
		Keyfile: "../etc/yourSite.key", // Your site private key
	})
	config := &tls.Config{}
	
	config.Certificates = make([]tls.Certificate, len(certs))
	for i, v := range certs  {
		config.Certificates[i], _ = tls.LoadX509KeyPair(v.CertFile, v.Keyfile)
	}
	conn, _ := net.Listen("tcp", ":8080")
	tlsListener := tls.NewListener(conn, config)
	httpsServer.Serve(tlsListener)
	fmt.Println("Listening on port 8080....")
}
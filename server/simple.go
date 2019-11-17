package main

import (
	"fmt"
	"log"
	"net/http"

	network "github.com/govishwavijay/learning/libs"
)

func main() {
	log.Println("Starting server________")
	network.RedirectToHTTPS(8080, "www.vishwavijay.com")

	server := http.Server{Addr: ":8443", Handler: nil}
	http.HandleFunc("/", handleRootContext)

	publicKey := "/home/vijay/public.cer"
	privateKey := "/home/vijay/privatekey.pem"
	if err := server.ListenAndServeTLS(publicKey, privateKey); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleRootContext(writer http.ResponseWriter, request *http.Request) {
	log.Println("Request came from : ", network.ReadRemoteIP(request), request.RequestURI)
	fmt.Fprintf(writer, "Hello there !! I am still working on this.")
}

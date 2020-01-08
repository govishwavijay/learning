package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	network "github.com/govishwavijay/learning/libs"
)

func main() {
	log.Println("Starting server________")
	network.RedirectToHTTPS(8080, "www.vishwavijay.com")

	network.SetStaticContentPath("/", network.GetHomeDirectory()+"static")

	log.Println("Will listen to port 8443....")
	server := http.Server{
		Addr:           ":8443",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        http.HandlerFunc(handleRootContext)}

	publicKey := network.GetHomeDirectory() + "/public.cer"
	privateKey := network.GetHomeDirectory() + "/privatekey.pem"
	if err := server.ListenAndServeTLS(publicKey, privateKey); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleRootContext(writer http.ResponseWriter, request *http.Request) {
	log.Println("Request came from : ", network.ReadRemoteIP(request), request.RequestURI)
	fmt.Fprintf(writer, "Hello there !! I am still working on this.")
}

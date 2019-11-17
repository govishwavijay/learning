package main

import (
	"fmt"
	"log"
	"net/http"

	network "github.com/govishwavijay/learning/libs"
)

func main() {
	log.Println("Fetching public ip address !")
	ip := network.FetchMyPublicIP()
	log.Println("Public IP address is ", ip)

	server := http.Server{Addr: ":8080", Handler: nil}
	http.HandleFunc("/hello", handleRootContext)
	http.HandleFunc("/favicon.ico", handleFavicon)
	network.SetStaticContentPath("/", network.GetHomeDirectory()+"/go/bin/static") //create static folder beside this and put UI content
	log.Println("Starting server.")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleRootContext(writer http.ResponseWriter, request *http.Request) {
	log.Println("Request came from : ", network.ReadRemoteIP(request), request.RequestURI)
	fmt.Fprintf(writer, "Hello there !!")
}

func handleFavicon(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "/home/vishwa/go/bin/static/favicon.ico")
}

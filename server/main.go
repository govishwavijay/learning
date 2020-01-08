package main

import (
	"fmt"
	"log"
	"net/http"

	network "github.com/govishwavijay/learning/libs"
)

func main() {
	http.HandleFunc("/", handleRootContext)
	err := http.ListenAndServeTLS(":8443", "/nas/server/root/private/public.cer", "/nas/server/root/private/privatekey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handleRootContext(writer http.ResponseWriter, request *http.Request) {
	log.Println("Request came from : ", network.ReadRemoteIP(request), request.RequestURI)
	fmt.Fprintf(writer, "Hello there !!")
}

func handleFavicon(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "/home/vishwa/go/bin/static/favicon.ico")
}

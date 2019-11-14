package network

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//FetchMyPublicIP : This function will fetch public ip from ifconfig.co
func FetchMyPublicIP(host ...string) string {
	url := "http://ifconfig.co/"
	if len(host) > 0 && len(host[0]) > 0 {
		url = host[0]
	}
	resp, err := http.Get(url)
	HandleError(err)
	body, err := ioutil.ReadAll(resp.Body)
	HandleError(err)
	resp.Body.Close()
	return strings.Trim(string(body), "\t \n")
}

//HandleError : This function will log the error
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//SetStaticContentPath : Set the path from where the static content will be served
func SetStaticContentPath(mapping string, path string) {
	log.Println("Static content will be loaded for mapping", mapping, "from path", path)
	fileSystem := http.FileServer(http.Dir(path))
	http.Handle(mapping, http.StripPrefix(mapping, fileSystem))
}

//ReadRemoteIP : Reads the remote ip address from where the request came.
func ReadRemoteIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

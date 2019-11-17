package network

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
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

//RedirectToHTTPS : This will redirect request comming to port 80 to 443
func RedirectToHTTPS(httpPort int, host string) {
	http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Request came to port 8080 : " + request.RequestURI)
		http.Redirect(writer, request, "https://"+host+request.RequestURI, 302)
	}))
}

//GetHomeDirectory : gets current user home directy
func GetHomeDirectory() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

package network

import (
	"io/ioutil"
	"log"
	"net/http"
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
	return string(body)
}

//HandleError : This function will log the error
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

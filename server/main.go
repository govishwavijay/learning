package main

import (
	"fmt"

	network "github.com/govishwavijay/libs"
)

func main() {
	ip := network.FetchMyPublicIP()
	fmt.Println("My IP address ba ", ip)
}
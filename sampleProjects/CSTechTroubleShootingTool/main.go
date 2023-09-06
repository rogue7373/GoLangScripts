package main

import (
	"fmt"

	"github.com/yourusername/yourpackagename/host"
)

func main() {
	hostname, err := host.GetHostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hostname:", hostname)
}

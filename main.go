package main

import(
	"fmt"
	"github.com/alexandernyquist/frogger"
)

func main() {
	proxy := frogger.Proxy{8082}
	err := proxy.Listen()
	if err != nil {
		fmt.Println("Could not listen on port 8082. Port probably already in use.")
	}
}
package main

import (
	"fmt"

	"net/http"

	"github.com/rishikeshKandi/Golang/main/pkg/handler"
)
	

	const portNumber = ":8090"
	

	func main() {
		http.HandleFunc("/", handler.Home)
		http.HandleFunc("/about", handler.About)
	

		fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
		_ = http.ListenAndServe(portNumber, nil)
	}

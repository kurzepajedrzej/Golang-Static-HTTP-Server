package main

import (
	"fmt"
	"net/http"
	"os"
)

var port string

func main() {
	static := http.FileServer(http.Dir("."))
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	http.Handle("/", static)
	fmt.Println("[*] - Server Listening on Port " + port)
	fmt.Println("[*] - http://localhost:" + port + "/")
	http.ListenAndServe(":"+port, nil)
}

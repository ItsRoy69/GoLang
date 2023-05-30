package main

import (
	"fmt"
	"net/http"
)

func helloWorldNinjaMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWorldNinjaMode")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello, World!</h1>")

}

func main() {

	server := http.Server{ //creating a server
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxNinjaMode http.ServeMux //creating a new server mux
	server.Handler = &muxNinjaMode
	muxNinjaMode.HandleFunc("/", helloWorldNinjaMode)

	server.ListenAndServe()
}

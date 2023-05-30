package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello, World!")
	case "/ninja":
		fmt.Fprintf(w, "Hello, Ninja!")
	default:
		fmt.Fprintf(w, "Error")
	}
	fmt.Printf("Handling request for %s\n", r.Method) //request function
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	w.Header().Set("Content-Type", "text/plain") //including the header tags as plain text
	fmt.Fprint(w, "<h1>Hello, World!</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) { //timeout function as timeout handler
	fmt.Println("Timeout Attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Did *not* timeout")
}

// func main() {
// 	http.HandleFunc("/", handleFunction)
// 	http.ListenAndServe("", nil)
// }

func main() {
	http.HandleFunc("/", htmlVsPlain) //htmlVsPlain
	http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)

	server := http.Server{ //creating a server
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	server.ListenAndServe()
}

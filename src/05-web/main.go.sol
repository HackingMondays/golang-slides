package main

import "log"
import "io"
import "net/http"

const port = ":8080"

func main() {
	log.Println("Server Running on", port)
	http.HandleFunc("/", myHandler)
	startServer()
}

func startServer() {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func myHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Got kicked at:", request.URL)
	io.WriteString(response, "Hello World")
}

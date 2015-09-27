package main

import "log"
import "io"
import "net/http"

const port = ":8080"

func main() {
	// TODO: afficher "Server Running on ..."
	// TODO: déclarer le handler avec http.HandleFunc( ... )
	startServer()
}

func startServer() {
	// TODO: listen and serve requests on port
	if err != nil {
		log.Fatal(err)
	}
}

func myHandler(response http.ResponseWriter, request *http.Request) {
	// TODO: afficher "Got kicked at:" + URL
	// TODO: écrire "Hello World" dans la réponse
}

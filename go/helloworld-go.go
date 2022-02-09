package main

// https://knative-v1.netlify.app/docs/serving/samples/hello-world/helloworld-go

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	provider := os.Getenv("CPLN_PROVIDER")
	location := os.Getenv("CPLN_LOCATION")

	log.Printf("helloworld: received a request. Provider: %s. Location: %s", provider, location)

	fmt.Fprintf(w, "<html>Hello World!<br/><br/>Provider: <b>%s</b><br/>Location: <b>%s</b></html>", provider, location)
}

func main() {
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Pipeline CI/CD 10/01/2022 v1"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}

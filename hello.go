package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/dharnitski/golang-codebuild/service"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

func main() {
	fmt.Println(service.Find("be"))
	handleRequests()
}

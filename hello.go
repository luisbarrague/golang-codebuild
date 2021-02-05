package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request){
	n, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		num, err2 := fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		fmt.Printf("%d bytes written, error found: %v.\n", num,err2)
	}
	fmt.Printf("%d bytes written.\n", n)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Printf("configured PORT: %s\n", os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		num, err2 := fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		fmt.Printf("%d bytes written, error found: %v.\n", num,err2)
	}
	handleRequests()
}

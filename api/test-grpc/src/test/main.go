package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Controller for the / route (home)
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page. Welcome!")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
    handleRequests()
}

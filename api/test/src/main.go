package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

// Controller for the / route (home)
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page. Welcome!")
}

func userName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page. Welcome!")
}

// Controller for the /items route
func returnAllItems(w http.ResponseWriter, r *http.Request) {
    // Code for returning all items here
}

// Controller for the /items/{id} route
func returnSingleItem(w http.ResponseWriter, r *http.Request) {
    // Code for returning a single item here
}

type Data1 struct {
	Title    string `json:"name"`
}
func createData(w http.ResponseWriter, r *http.Request){
    var data1 = Data1{}
    data1.Title = "string"

    // jsonエンコード
    outputJson, err := json.Marshal(&data1)
    if err != nil {
        panic(err)
    }

    // jsonヘッダーを出力
    w.Header().Set("Content-Type", "application/json")

    // jsonデータを出力
    fmt.Fprint(w, string(outputJson))
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/user/get", createData)
    myRouter.HandleFunc("/items", returnAllItems)
    myRouter.HandleFunc("/items/{id}", returnSingleItem)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
    handleRequests()
}

package main

import (
    "net/http"
    "fmt"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got:", r.URL.Query());
    fmt.Fprint(w, "Hello World")
}

func main() {
    http.HandleFunc("/", mainPage)
    http.ListenAndServe(":8080", nil)
}

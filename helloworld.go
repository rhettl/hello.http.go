package main

import (
    "html/template"
    "net/http"
    "fmt"
)

var index = template.Must(template.ParseFiles(
    "templates/_base.html",
    "templates/index.html",
))

func hello(w http.ResponseWriter, req *http.Request) {
    if err := index.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Starting Server...");
    if err := http.ListenAndServe(":8000", nil); err != nil {
        panic(err)
    }
}

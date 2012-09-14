package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "HELLO WORLD FROM A LISTENER")
}

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Starting Server...");
    if err := http.ListenAndServe(":8000", nil); err != nil {
        panic(err)
    }
}

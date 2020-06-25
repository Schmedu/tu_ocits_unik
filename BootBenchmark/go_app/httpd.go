package main

import (
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("http://127.0.0.1:5000/stop/gp_App1")
    fmt.Printf("being busy...")
    time.Sleep(5 * time.Second)
    fmt.Printf("..bye")
//     http.HandleFunc("/", handler)
//     http.ListenAndServe(":8080", nil)
}

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "my first unikernel!")
// }

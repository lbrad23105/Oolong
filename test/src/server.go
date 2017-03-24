package main

import "net/http"
import "fmt"

func main() {
   mux := http.NewServeMux()
   files := http.FileServer(http.Dir(config.Static))
   mux.Handle("/static/",http.StripPrefix("/static",files))   mux.HandleFunc("/", index)
   http.ListenAndServe(":5000", mux)
}
package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "<h1>Hello</h1>")
}

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

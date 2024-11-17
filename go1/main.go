package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server1です"))
	})

	http.ListenAndServe(":8080", nil)
}

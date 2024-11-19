package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1秒待つ
		time.Sleep(1 * time.Second)
		w.Write([]byte("server2です"))
	})

	http.ListenAndServe(":8080", nil)
}

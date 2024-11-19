package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 10秒待つ
		time.Sleep(10 * time.Second)
		w.Write([]byte("server3です"))
	})

	http.ListenAndServe(":8080", nil)
}

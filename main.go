package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		w.Header().Add("Content-Type", "application/json")
		_ = encoder.Encode(map[string]string{
			"ok": "true",
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))

}

func main() {
	http.HandleFunc("/", rootHandler)

	println(fmt.Sprintf("Listening on port %d", 8080))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println(fmt.Sprintf("Error starting server %v", err))
	}
}

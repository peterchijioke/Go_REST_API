package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", test)
	http.ListenAndServe(":5000", r)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is test"))
	// json.NewDecoder(w).Encode()
}

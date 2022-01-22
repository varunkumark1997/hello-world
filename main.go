package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homepage)

	http.ListenAndServe(":80", mux)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

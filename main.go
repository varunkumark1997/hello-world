package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", homepage)
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {

	html := `
<h1>Hello World!</h1>
<div class="container">
  <canvas id="c" width="650" height="650" style="background:#000"></canvas>
</div>`
	io.WriteString(w, html)
}

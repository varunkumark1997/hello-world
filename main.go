package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", homepage)
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {

	size := 650
	sizeString := strconv.Itoa(size)
	randx := rand.Int() % size
	randy := rand.Int() % size

	html := `
<h2>Snake Game!!</h1>

<div class="container">
  <canvas id="myCanvas" width=` + sizeString + ` height=` + sizeString + ` style="background:#000"></canvas>
</div>

<script>
	var canvas = document.getElementById("myCanvas");
	var ctx = canvas.getContext("2d");
	ctx.fillStyle = 'white';
	ctx.fillRect(` + strconv.Itoa(randx) + `, ` + strconv.Itoa(randy) + `, 10, 10);  
</script>`
	io.WriteString(w, html)
}

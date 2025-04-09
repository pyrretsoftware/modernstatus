package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/axist.css", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ba, err := os.ReadFile("./src/status.html")
		if (err != nil) {return}

		page := strings.ReplaceAll(string(ba), "$${{StatusJson}}", r.Header.Get("x-ray-plugin-data"))

		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(page))
	})

	port := os.Getenv("RAY_PORT")
	if (port != "") {
		http.ListenAndServe(":" + port, nil)
	} else {
		log.Fatal("what (not using ray???)")
	}
}
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"quintus"
)

func checkerror(err error) {
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}
}

func handle_index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/layout.template", "templates/index.template")

	checkerror(err)

	t.Execute(w, someq)
}

var someq []quintus.Quintus

func init() {
	qs := make([]quintus.Quintus, 1)

	qs[0].Name = "Go"
	qs[0].Date = "Today"
	qs[0].Times = 0

	quintus.Write(qs)

	someq = quintus.Read()
	log.Println("[INFO]", "LOADED GOB", someq)
}

func main() {
	workingdirectory, err := os.Getwd()

	checkerror(err)

	log.Println("[INFO]", "Starting server in directory", workingdirectory)

	http.HandleFunc("/", handle_index)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(workingdirectory+`/public`))))

	err = http.ListenAndServe(":9090", nil)

	checkerror(err)
}

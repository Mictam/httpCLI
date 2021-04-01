package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler( w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("static/basictemplate.html")
	t.Execute(w, p)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>We can GO!</h1>")
}

func help() {

}

func versionInfo() {

}

func runServer(htmlFile string) {

}
func main() {
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/agg/", newsAggHandler)
	//http.ListenAndServe(":8080", nil)
	switch os.Args[0] {
		case "help":
			help()
		case "version":
			versionInfo()
		case "run":
			if os.Args[1] != "--file" || os.Args[2] == "" {
				fmt.Println("Invalid argument provided, use \"run --file <index.html>\"")
				return
			}
			runServer(os.Args[2])
		default:
			fmt.Println("Invalid argument provided, use \"httpCLI help\" to get more info")
			return
	}
}
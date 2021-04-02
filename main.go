package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	AppVersion string = "v0.1"
	HtmlPath string = "./static/basictemplate.html"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HtmlPath)
}

func help() {
	fmt.Print("\nThis is simple CLI app that runs HTTP server at localhost:8080 \n\nExample usage:" +
		"\n  httpCLI run --file <index.html>\n\nAvailable " +
		"Commands:\n  version                       Displays app version info\n  " +
		"help                          Help about any command\n  " +
		"run [--file <index.html>]     Runs http server that serves alternate html file (file must be in current catalog or filepath)\n")
}

func versionInfo() {
	fmt.Print("\n App version is " + AppVersion +"\n")
}

func runServer(HtmlFile string) {
	HtmlPath = HtmlFile
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	switch os.Args[1] {
		case "help":
			help()
		case "version":
			versionInfo()
		case "run":
			if len(os.Args) == 2 {
				runServer(HtmlPath)
				return
			}
			if os.Args[2] == "--file" && len(os.Args) < 4 {
				fmt.Println("Invalid argument provided, use \"run --file <index.html>\"")
				return
			}
			runServer(os.Args[3])
		default:
			fmt.Println("Invalid argument " + os.Args[1] + " provided, use \"httpCLI help\" to get more info")
			return
	}
}
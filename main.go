package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, _ := ResourceFile("index.html")
	fmt.Fprint(w, "<!-- test -->")
	fmt.Fprint(w, string(file))
}

func GetStylesheet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	file, _ := ResourceFile("styles.css")
	fmt.Fprint(w, string(file))
}

func GetScript(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	file, _ := ResourceFile("scripts.js")
	fmt.Fprint(w, string(file))
}

// at the moment it's not possible to load raw templates. W
func GetTemplate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	file, _ := ResourceFile("templates.tpl.js")
	fmt.Fprint(w, string(file))
}

func GetBrand(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	file, _ := ResourceFile("brand.svg")
	fmt.Fprint(w, string(file))
}

func main() {

	var port, seconds int
	var filename string

	flag.IntVar(&port, "port", 54321, "Port to listen to.")
	flag.StringVar(&filename, "file", "./talkapply.json", "Backup file. (see: seconds)")
	flag.IntVar(&seconds, "seconds", 5, "Every x seconds a backup of all in memory data is created on disk.")
	flag.Parse()

	initStorage(seconds, filename)

	router := httprouter.New()
	router.GET("/", GetIndex)
	router.GET("/styles.css", GetStylesheet)
	router.GET("/scripts.js", GetScript)
	router.GET("/templates.tpl.js", GetTemplate)
	router.GET("/brand.svg", GetBrand)

	fmt.Printf("\nListening on port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

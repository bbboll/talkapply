package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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

func GetBrand(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	file, _ := ResourceFile("brand.svg")
	fmt.Fprint(w, string(file))
}


func main() {

	router := httprouter.New()
	router.GET("/", GetIndex)
	router.GET("/styles.css", GetStylesheet)
	router.GET("/scripts.js", GetScript)
	router.GET("/brand.svg", GetBrand)

	log.Fatal(http.ListenAndServe(":54321", router))
}
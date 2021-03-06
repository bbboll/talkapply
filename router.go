package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
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

func initRouter(port int) {

	router := httprouter.New()
	router.GET("/", GetIndex)
	router.GET("/styles.css", GetStylesheet)
	router.GET("/scripts.js", GetScript)
	router.GET("/templates.tpl.js", GetTemplate)
	router.GET("/brand.svg", GetBrand)

	fmt.Printf("\nListening on port %d\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(150)
	}
}

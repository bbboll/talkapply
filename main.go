package main

import (
	"flag"
)

func main() {

	var port, seconds int
	var filename string

	flag.IntVar(&port, "port", 54321, "Port to listen to.")
	flag.StringVar(&filename, "file", "./talkapply.json", "Backup file. (see: seconds)")
	flag.IntVar(&seconds, "seconds", 5, "Every x seconds a backup of all in memory data is created on disk.")
	flag.Parse()

	initRouter(port)
	initStorage(seconds, filename)
}

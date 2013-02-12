package main

import (
	"net/http"
	"fmt"
	"flag"
	"log"
)

var (
	port string
	public string
	help bool
)

func init() {
	flag.StringVar(&port,"port",":8080","Selects the port you wish to use")
	flag.StringVar(&public,"public","./","Serve all content of directory")
	flag.BoolVar(&help,"help",false,"Display this help message")
}

func main() {
	flag.Parse()
	if help == true {
		flag.Usage()
	}
	http.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir(public))))
	fmt.Printf("Listening on %s, serving public from \"%s\"...\n",port,public)
	log.Fatal(http.ListenAndServe(port,nil))
}

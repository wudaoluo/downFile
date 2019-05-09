package main

import (
	"log"
	"net/http"
	"os"
)


func main() {
	var port = "80"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	http.Handle("/",http.FileServer(http.Dir("/tmp")))
	log.Println("downFile server listen port",port)
	err  := http.ListenAndServe(":"+port,nil)
	if err!=nil{
		panic(err)
	}
}

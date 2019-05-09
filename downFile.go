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

	h := http.FileServer(http.Dir("./data"))
	log.Println("downFile server listen port",port)
	err := http.ListenAndServe(":"+port, h)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

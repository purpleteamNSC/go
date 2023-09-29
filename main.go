package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fs:=http.FileServer(http.Dir("/workspaces/go/share"))
	fmt.Println("HttpServer Go")
	log.Fatal(http.ListenAndServe(":9000",fs))
}
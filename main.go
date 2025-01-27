package main

import (
	"fmt"
	"go-server/routes"
	"log"
	"net/http"
)

func main(){
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", routes.FormHandler)

	fmt.Println("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
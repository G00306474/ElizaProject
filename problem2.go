package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Templatedata struct{
	Message string
	
}

func templateHandler(w http.ResponseWriter, r *http.Request){
	 t, _ := template.ParseFiles("webFolder/about.html")
	 t.Execute(w, Templatedata {Message: "Eliza"} )
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./webFolder")))
	fmt.Println("listen to Eliza on 3000")
	http.HandleFunc("/guess", templateHandler)

    http.ListenAndServe(":3000", nil)
}
package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/template", func(writer http.ResponseWriter, request *http.Request) {

		parse, _ := template.New("hello.html").ParseFiles("./files/hello.html")

		err := parse.Execute(writer, gin.H{"name": "zzh"})
		if err != nil {
			return
		}

	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}

}

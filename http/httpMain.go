package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("./files/hello.html")
	_, err := fmt.Fprintf(w, string(file))
	if err != nil {
		return
	}

}

func main() {

	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("http serve fail", err)
	}
}

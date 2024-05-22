package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", req.PostForm)
}

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

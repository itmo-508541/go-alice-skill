package main

import (
	"net/http"
)

func main() {
	var h MyHandler
	err := http.ListenAndServe(`:8080`, h)
	if err != nil {
		panic(err)
	}
}

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

type MyHandler struct{}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Привет!")
	res.Write(data)
}

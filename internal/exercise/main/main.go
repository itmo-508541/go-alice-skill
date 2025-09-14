package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Привет!<br />"))

	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainPage)
	mux.HandleFunc(`/api`, apiPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

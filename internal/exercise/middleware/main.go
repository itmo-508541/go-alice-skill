package main

import (
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет"))
}

func MiddlewareHandler(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// здесь пишем логику обработки
		// например, разрешаем запросы cross-domain
		// ...
		// замыкание: используем ServeHTTP следующего хендлера
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(`/`, MiddlewareHandler(http.HandlerFunc(rootHandle)))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

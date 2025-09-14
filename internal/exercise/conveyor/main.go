package main

import (
	"net/http"
	"time"
)

type MiddlewareFunc func(http.Handler) http.Handler

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[rootHandle]"))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/favicon.ico")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://yandex.ru/", http.StatusMovedPermanently)
}

func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// добавить заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")

		w.Write([]byte("[HeadersMiddleware]"))
		next.ServeHTTP(w, r)
	})
}

func AccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// сделать ещё
		hasAccess := true
		if hasAccess {
			w.Write([]byte("[AccessMiddleware]"))
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func Conveyor(h http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func main() {
	сonveyor := Conveyor(http.HandlerFunc(rootHandle), HeadersMiddleware, AccessMiddleware)

	mux := http.NewServeMux()
	mux.Handle(`/`, http.TimeoutHandler(сonveyor, 10*time.Second, "Service Unavailable"))
	mux.Handle(`/dummy`, http.RedirectHandler("https://google.com", http.StatusMovedPermanently))
	mux.Handle(`/search/`, http.HandlerFunc(redirect))

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/favicon.ico", faviconHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

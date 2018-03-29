package main

import (
	"weChat/config"
	"net/http"
	"fmt"
	"time"
	"html"
)

func main() {
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", config.Port()),
		Handler:        http.HandlerFunc(ServeHttp),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}

	server.ListenAndServe()
}

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
}

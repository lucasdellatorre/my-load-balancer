package main

import (
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"
)

var count = 0

func hub(w http.ResponseWriter, r *http.Request) {
	routePath := html.EscapeString(r.URL.Path) // EscapeString prevent injection vulnerabilities
	logRequestDetails(w, r)
	count++
	fmt.Println(count)
	switch routePath {
	case "/":
		fmt.Fprintf(w, "Hello from backend server")
	case "/foo":
		fmt.Println("hit foo")
	default:
		http.NotFound(w, r)
	}
}

func logRequestDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received request from %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host: %s\n", strings.Split(r.Host, ":")[0])
	fmt.Fprintf(w, "User-Agent: %s\n", r.UserAgent())
	fmt.Fprintf(w, "Accept: %s", r.Header.Get("Accept"))
}

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(hub),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Backend is running...")

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

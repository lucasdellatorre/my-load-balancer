package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"
)

func hub(w http.ResponseWriter, r *http.Request) {
	routePath := html.EscapeString(r.URL.Path) // EscapeString prevent injection vulnerabilities
	logRequestDetails(w, r)
	switch routePath {
	case "/":
		req, err := http.NewRequest("GET", "http://localhost:8080", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		fmt.Fprintf(w, "Response from server: %s %d %s", resp.Proto, resp.StatusCode, resp.Status)
		io.WriteString(w, string(body))
	default:
		http.NotFound(w, r)
	}
}

func logRequestDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received request from %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host: %s\n", strings.Split(r.Host, ":")[0])
	fmt.Fprintf(w, "User-Agent: %s\n", r.UserAgent())
	fmt.Fprintf(w, "Accept: %s\n", r.Header.Get("Accept"))
}

func main() {
	s := &http.Server{
		Addr:           ":80",
		Handler:        http.HandlerFunc(hub),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)

	// go func() {
	fmt.Println("Load balancer is running")
	if err := s.ListenAndServe(); err != nil && http.ErrServerClosed != err {
		panic(err)
	}
	// }()

	// 	fmt.Println("LoadBalancer is running...")

	// 	<-stop

	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	fmt.Println("Shutdown...")

	// 	if err := s.Shutdown(ctx); err != nil {
	// 		panic(err)
	// 	}

	// fmt.Println("Server stopped")
}

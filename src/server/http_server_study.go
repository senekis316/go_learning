package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, " "))
	}
	io.WriteString(w, "ok")
}

func GetRealIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-IP")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarder-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// call the original http.Handler we're wrapping
		h.ServeHTTP(w, r)
		fmt.Println("ClientIP:", GetRealIP(r))
		fmt.Println("HTTP Status Code:", 200)
	}

	// http.HandlerFunc wraps a function so that it
	// implements http.Handler interface
	return http.HandlerFunc(fn)
}

func main() {

	os.Setenv("VERSION", "1.0")

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8000", logRequestHandler(mux))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Http服务器开始监听")

}

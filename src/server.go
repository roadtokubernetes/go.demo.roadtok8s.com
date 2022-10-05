package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	var port string
	port = getEnv("PORT", "8012")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/healthz", Healthz)
	fmt.Println("Running from http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

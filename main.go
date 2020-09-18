package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	portNumber := os.Getenv("CICD_PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World to CI/CD"))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", portNumber), nil)
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

// Report what you are
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Cloud! %s", os.Args[1:])
}

func main() {
	http.HandleFunc("/", handler)     // redirect all urls to the handler function
	http.ListenAndServe(":9999", nil) // listen for connections at port 9999 on the local machine
}

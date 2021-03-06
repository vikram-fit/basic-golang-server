package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "<h1> Welcome to my site </h1>")
	} else if r.URL.Path == "/contact" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:vikramt@fitternity.com\">Vikram T</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> 404 Page Not Found </h1>")
	}
	// fmt.Println("New request has been made")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
	// fmt.Println("Hello World!!")
}

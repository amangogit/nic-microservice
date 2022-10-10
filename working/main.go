package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Println("Hello World!")
		d, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(resp, "Oops", http.StatusBadRequest)
			// resp.WriteHeader(http.StatusBadRequest)
			// resp.Write([]byte("Oops"))
			return
		}

		log.Printf("Data %s\n", d)

		fmt.Fprintf(resp, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World!")
	})

	// bind address
	http.ListenAndServe(":9990", nil)
}

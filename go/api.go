package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"net/http"
)

type numbers struct {
	value []int
}

func main() {
	primes := numbers{[]int{2, 3, 5, 7, 11, 13}}
	out, _ := xml.Marshal(primes)
	fmt.Println(string(out))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	"github.com/jchannon/negotiator"
	"github.com/julienschmidt/httprouter"
)

type numbers struct {
	Value []int `xml:"value"`
}

type sum struct {
	Sum int64 `xml:",chardata" json:"sum"`
}

func main() {
	router := httprouter.New()
	router.GET("/add/:x/to/:y", func(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
		x, _ := strconv.ParseInt(params.ByName("x"), 10, 64)
		y, _ := strconv.ParseInt(params.ByName("y"), 10, 64)

		if err := negotiator.Negotiate(responseWriter, request, sum{x + y}); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
	})

	router.POST("/add", func(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
		decoder := xml.NewDecoder(request.Body)
		var payload numbers
		err := decoder.Decode(&payload)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()
		var total int64

		for val := range payload.Value {
			total += int64(val)
		}

		if err := negotiator.Negotiate(responseWriter, request, sum{total}); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

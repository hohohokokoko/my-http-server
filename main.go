// main function for myhttp

package main

import (
	"log"
	"myhttp/myhttp"
)

func main() {
	r := myhttp.NewRouter()

	helloHandler := func(resp *myhttp.Response, req *myhttp.Request) {
		resp.RawText(200, "Hello world!")
	}

	r.GET("/hello", helloHandler)

	log.Fatal(myhttp.ListenAndServe(":8000", r))
}

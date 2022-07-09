package main

import (
	"fmt"
	"net/http"

	"github.com/ogwurujohnson/loadbalancer/lib"
)

func main() {
	servers := []lib.Server{
		lib.NewSimpleServer("https://www.google.com"),
		lib.NewSimpleServer("https://www.bing.com"),
		lib.NewSimpleServer("https://www.duckduckgo.com"),
	}

	lb := lib.NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}

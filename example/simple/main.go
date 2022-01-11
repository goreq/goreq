package main

import (
	"fmt"

	"github.com/hadihammurabi/gore"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	g := gore.New(
		gore.WithBaseURL("https://my-json-server.typicode.com/hadihammurabi/flutter-webservice"),
	)

	resp, err := g.Get("/contacts", nil)
	panicOnError(err)
	defer resp.Body.Close()

	var data interface{}
	err = resp.Json(&data)
	panicOnError(err)

	fmt.Println(data)
}

package main

import (
	"fmt"
	"sync"

	"github.com/goreq/goreq"
)

type Contact map[string]interface{}
type Contacts []Contact

func getContacts(g goreq.Gore) Contacts {
	contacts := make(Contacts, 0)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(g, fmt.Sprint(i), &contacts, &wg)
	}

	wg.Wait()

	return contacts
}

func worker(g goreq.Gore, id string, contacts *Contacts, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := g.Get("/contacts/"+id, nil)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var result Contact
	err = resp.Json(&result)
	if err != nil {
		return
	}

	*contacts = append(*contacts, result)
}

func main() {
	g := goreq.New(
		goreq.WithBaseURL("https://my-json-server.typicode.com/hadihammurabi/flutter-webservice"),
	)

	contacts := getContacts(g)
	fmt.Println(contacts)
}

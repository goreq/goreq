# GoReq (Go Requester)
Simple HTTP client for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/goreq/goreq)](https://goreportcard.com/report/github.com/goreq/goreq)


* [Documentation](https://pkg.go.dev/github.com/goreq/goreq)
* [Coverage](https://gocover.io/github.com/goreq/goreq)
* [Examples](https://github.com/goreq/example)

# Example

```go
resp, err := goreq.Get("https://api.products.com/entities")
if err != nil {
  panic(err)
}

defer resp.Body.Close()

// do anything with the resp object
```

or, using base URL like this

```go
g := goreq.New(
  gore.WithBaseURL("https://api.products.com"),
)

resp, err := g.Get("/entities")
if err != nil {
  panic(err)
}

defer resp.Body.Close()

// do anything with the resp object
```

> more examples, see [https://github.com/goreq/example](https://github.com/goreq/example)

# Key Concept
* Reusable, prevent options rewrite using single object for multiple request
* Flexible, because it use raw response object

# Contributing
Let's code together by following [Contributing Guidelines](./CONTRIBUTING.md)

# License
This library under the [GNU General Public License v3.0](./LICENSE)

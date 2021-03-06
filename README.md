# GoReq (Go Requester)
Simple HTTP client for Go

# Example

```go
resp, err := goreq.Get("https://api.products.com/entities", nil)
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

resp, err := g.Get("/entities", nil)
if err != nil {
  panic(err)
}

defer resp.Body.Close()

// do anything with the resp object
```

# Key Concept
* Reusable, prevent options rewrite using single object for multiple request
* Flexible, because it use raw response object

# Stargazers over time

[![Stargazers over time](https://starchart.cc/hadihammurabi/gore.svg)](https://starchart.cc/hadihammurabi/gore)


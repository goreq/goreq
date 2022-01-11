# GoRe (Go Requester)
Simple HTTP client for Go

# Example

```go
g := gore.New(
  gore.WithBaseURL("https://api.products.com"),
)

resp, err := g.Get("/entities", nil)
if err != nil {
  panic(err)
}

defer resp.Body.Close()

// do anything with the resp object
```

# Features
* Reusable, prevent options rewrite using single object for multiple request
* Flexible, because it use raw response object

# Todos
* [x] Global error handler
* [x] Global request and response hooks
* [x] Use custom request object

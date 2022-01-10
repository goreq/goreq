# GoRe (Go Requester)
Simple HTTP client for Go

# Example

```go
g := gore.New(
  gore.WithBaseURL("https://api.products.com"),
)

resp, err := g.Get("/entities")
if err != nil {
  panic(err)
}

defer resp.Body.Close()

// do anything with the resp object
```

# Features
* Reusable, prevent rewrite options with single object for multiple request
* Flexible, because it use raw response object

# Todos
* [ ] Global error handler
* [ ] Global request and response hooks

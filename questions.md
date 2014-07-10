- What's the point of the error type in Go? It's pretty much just used to Print using fmt.

- http://golang.org/pkg/errors/
- http://www.golang-book.com/13/index.htm#section4

```go
err := errors.New("error message")
    if err != nil {
        fmt.Println(err)
    }
```

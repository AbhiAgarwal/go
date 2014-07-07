# Random notes

- argument to go/defer must be function call

```go
defer func() {
    fmt.Println("x")
}
```

Doesn't work! You've to pass in a function

```go

m := func() {
    fmt.Println("x")
}

defer m()
```

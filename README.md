Notes for some presentations on Go to beginners. 

# Observations

```go
if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
}
```

or 

```go
js := []byte{}
js, err = json.Marshal(currentBox)
if err != nil {
    panic(err)
}
```

# Go-specific things

- Slices
- make/delete/new function
- Variadic function
    - takes zero or more parameters in args
- unpacking slices/array into Variadic functions
    - callFunction(sliceName...)
- Closures - function declaration inside function
- Defer, panic, and recover
- Garbage collected
- Struct methods, and receivers
- Embedded Types, and anonymous fields
- Interfaces, and “method set[s]”
- Concurrency, and Goroutines and channels
    - Select (like switch) statement for channel
    - Buffered Channels
- Error Handling in Go is IMPORTANT
    - Error type
- Declaring your own type `type Fame []int`

# Things I don't understand well yet

- Interfaces, and “method set[s]”

# Reading

- [Go for the paranoid network programmer](http://www.slideshare.net/feyeleanor/go-for-the-paranoid-network-programmer)
- [Build Web Application with Golang](https://github.com/astaxie/build-web-application-with-golang/tree/master/en)
- [Golang Removing fields from struct or hiding them in JSON Response](http://stackoverflow.com/questions/17306358/golang-removing-fields-from-struct-or-hiding-them-in-json-response)
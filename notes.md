# Random notes

- Concurrency Is Not Parallelism

    - Concurrency is about dealing with a lot of things at once
    - Parallelism is about doing a lot of things at once
    - Parallelism comes in when you have 2 gophers and you want them to move at the same time. 
    - You can be concurrent by having 2 gophers do different tasks, but you can establish a channel between them and make them communicate.
    - It's like having a single core computer.
    - Parallelism is having 2 tasks execute simultaneously, not just having 2 things
        - which is what can happen in concurrency.

- Hashes are frequently used in programming for everything from looking up data to easily detecting changes.

- _ (blank) is a special name of variable, any value that is given to it will be ignored.

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

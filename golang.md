# golang

Notes on the Go programming language.

Created on Jan 1, 2023.


## Concepts

- Context
    - e.g.: timeouts
- Defer (`defer`)
    - function call executed at the end of the enclosing function
    - usually used for cleanup
- Closures
    - just anonymous, inline functions
- Go routines (`go func`)
    - similar to threads, but managed by the Go runtime
- `sync.WaitGroup`s
    - to wait for multiple Go routines to finish
- `sync.Mutex` (mutual exclusion)
    - read/write lock to prevent racing conditions
- Channels (`chan`)
    - to _communicate_ data between Go routines (not shared!)
        - data are copied; don't pass pointers!
    - can be buffered or unbuffered
        - unbuffered channel (`make(chan int)`, with no capacity) needs to emit and receive simultaneously
        - buffered channel (e.g.: `make(chan int, 5)`) does not block (deadlock) until buffer is full


## Patterns

### Generators

=> "functions that return the next value in a sequence each time the function is called."

See: [link](http://www.golangpatterns.info/concurrency/generators)


### Futures (async/awit)

=> To compute a value in another process before you need to use it.

```go
future := make(chan int, 1)
go func() { future <- process() }() // async
result := <-future                  // await
```


### Scatter/Gather

- Great to use when the number of units of work is precisely known.
    - e.g.: fetch 10 images from an API

```go
// Scatter
c := make(chan result, 10)
for i := 0; i < cap(c); i++ {
    go func() {
        val, err := process()
        c <- result{val: val, err: err}
    }()
}

// Gather
var total int
for i := 0; i < cap(c); i++ {
    result := <-c
    if result.err == nil {
        total += result.val
    }
}
```


## Resources

- Go by Example [link](https://gobyexample.com/)
- "GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming" [video](https://www.youtube.com/watch?v=PTE4VJIdHPg)
- "GothamGo 2018 - Things in Go I Never Use by Mat Ryer" [video](https://www.youtube.com/watch?v=5DVV36uqQ4E)



- https://www.youtube.com/watch?v=VkGQFFl66X4



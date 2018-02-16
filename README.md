A Simple stop watch written in Go.

> Inspired by Spring framework!

## Installation

    go get -u github.com/nekolunar/stopwatch

## Usage

```go
sw := stopwatch.New() // or stopwatch.NewID("my test")

sw.Start()
// do some work.
time.Sleep(1 * time.Second)
sw.Stop()

log.Println("elapsed", sw.Elapsed())

sw.StartTask("my task 1")
time.Sleep(150 * time.Millisecond)
sw.Stop()

sw.StartTask("my task 2")
time.Sleep(150 * time.Millisecond)
sw.Stop()

log.Println(sw)
```

more simply

```go
start := stopwatch.Start()

// do some work.
time.Sleep(1 * time.Second)

elapsed := stopwatch.Stop(start)

log.Println("elapsed", elapsed.Seconds())
```

### TODOs:

- test/lint
- godoc

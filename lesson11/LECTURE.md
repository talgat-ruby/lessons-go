# Lesson 11: Concurrency (CSP)

**Concurrency**:

Concurrency is a concept related to the structure and organization of a program.
It’s about managing multiple tasks and their execution, often at the same time, without necessarily requiring them to
run simultaneously. Concurrency can be thought of as a program’s ability to execute multiple tasks in overlapping
time frames, without guaranteeing that they will execute simultaneously.

> **Real-world analogy**: Imagine a restaurant kitchen with multiple chefs.
> Each chef is working on their assigned tasks, like chopping vegetables, grilling meat, and preparing sauces.
> While these chefs work independently on their tasks, the restaurant manager coordinates their work,
> ensuring that dishes are prepared efficiently and ready to be served.
> This is analogous to concurrent programming, where tasks are managed efficiently,
> and the CPU switches between them to ensure progress.

**Parallelism**:

Parallelism, on the other hand, is the simultaneous execution of multiple tasks or processes,
typically to improve performance by taking advantage of multiple CPU cores or processors.
It’s about running multiple threads or processes in parallel to achieve a speedup in execution time.

> **Real-world analogy**: Think of an assembly line in a car manufacturing plant.
> Each worker is responsible for a specific part of the car assembly, and these workers are all working simultaneously.
> This is analogous to parallel programming,
> where multiple threads or processes are actively executing at the same time to achieve faster computation.

Differences:

1. Concurrency is about managing and organizing tasks,
   while parallelism is about executing tasks simultaneously to improve performance.

2. Concurrency doesn't guarantee that tasks will run in true parallel;
   they may share the same resources or take turns using them.
   Parallelism, on the other hand, ensures tasks are actively running at the same time on separate resources.

[More](https://www.baeldung.com/cs/concurrency-vs-parallelism)

## Goroutines

The goroutine is the core concept in Go’s concurrency model. To understand goroutines, let’s define a couple of terms.
The first is **process**. A **process** is an instance of a program that’s being run by a computer’s operating system.
The operating system associates some resources, such as memory, with the process and makes sure that other processes
can’t access them. A **process** is composed of one or more **threads**.
A **thread** is a unit of execution that is given some time to run by the operating system.
**Threads** within a process share access to resources.
A **CPU** can execute instructions from one or more threads at the same time, depending on the number of cores.
One of the jobs of an operating system is to schedule threads on the CPU to make sure that every
**process** (and every thread within a process) gets a chance to run.

Goroutines are lightweight, independently scheduled functions that can be executed concurrently.
They are similar to threads but are managed by Go’s runtime, making them more efficient.
Goroutines enable concurrent execution in Go.

- Goroutine creation is faster than thread creation, because you aren’t creating an operating system–level resource.
- Goroutine initial stack sizes are smaller than thread stack sizes and can grow as needed.
  This makes goroutines more memory efficient.
- Switching between goroutines is faster than switching between threads because it happens entirely within the process,
  avoiding operating system calls that are (relatively) slow.
- The goroutine scheduler is able to optimize its decisions because it is part of the Go process.
  The scheduler works with the network poller,
  detecting when a goroutine can be unscheduled because it is blocking on I/O.
  It also integrates with the garbage collector,
  making sure that work is properly balanced across all the operating system threads assigned to your Go process.

## Channels

Channels are a communication mechanism in Go that allows goroutines to communicate and synchronize.
They facilitate coordination between concurrent tasks, which is crucial for many concurrent programs.

- Create

```go
ch := make(chan int) // unbuffered
bch := make(chan int 1) // buffered
```

- Read

```go
a, ok := <-ch // closed will return value, false
```

- Write

```go
ch <- b
```

- Iterate

```go
for v := range ch {
fmt.Println(v)
}
```

- Close

```go
close(ch)
```

|       | Unbuffered, open                 | Unbuffered, closed                                | Buffered, open                      | Buffered, closed                                                                                                  | Nil          |
|-------|:---------------------------------|---------------------------------------------------|-------------------------------------|-------------------------------------------------------------------------------------------------------------------|--------------|
| Read  | Pause until something is written | Return zero value (use comma ok to see if closed) | Pause if buffer is empty            | Return a remaining value in the buffer; if the buffer is empty, return zero value (use comma ok to see if closed) | Hang forever |
| Write | Pause until something is read    | **PANIC**                                         | Pause if buffer is full             | **PANIC**                                                                                                         | Hang forever |
| Close | Works                            | **PANIC**                                         | Works, remaining values still there | **PANIC**                                                                                                         | **PANIC**    |

## `select`

The select statement is the other thing that sets apart Go’s concurrency model.
It is the control structure for concurrency in Go, and it elegantly solves a common problem:
if you can perform two concurrent operations, which one do you do first?
You can’t favor one operation over others, or you’ll never process some cases. This is called _starvation_.

```go
select {
case v := <-ch:
fmt.Println(v)
case v := <-ch2:
fmt.Println(v)
case ch3 <- x:
fmt.Println("wrote", x)
case <-ch4:
fmt.Println("got value on ch4, but ignored it")
}
```

What happens if multiple cases have channels that can be read or written?
The `select` algorithm is simple: it picks randomly from any of its cases that can go forward;
order is unimportant. This is very different from a `switch` statement,
which always chooses the first case that resolves to true.
It also cleanly resolves the _starvation_ problem,
as no case is favored over another and all are checked at the same time.

### deadlocks

If you have two goroutines that both access the same two channels,
they must be accessed in the same order in both goroutines, or they will _deadlock_.
This means that neither one can proceed because they are waiting on each other.
If every goroutine in your Go application is deadlocked, the Go runtime kills your program.

```go
func main() {
ch1 := make(chan int)
ch2 := make(chan int)
go func () {
inGoroutine := 1
ch1 <- inGoroutine
fromMain := <-ch2
fmt.Println("goroutine:", inGoroutine, fromMain)
}()
inMain := 2
ch2 <- inMain
fromGoroutine := <-ch1
fmt.Println("main:", inMain, fromGoroutine)
}
```

## for loops

The for range form of the **for loop** can be used to receive values from a channel until it is closed.

```go
package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
```

_example2_

## Best Practices

### Clean up

Whenever you launch a goroutine function, you must make sure that it will eventually exit.
Unlike variables, the Go runtime can’t detect that a goroutine will never be used again.
If a goroutine doesn't exit, all the memory allocated for variables on its stack remains allocated and
any memory on the heap that is rooted in the goroutine’s stack variables cannot be garbage collected.
This is called a _goroutine leak_.

### Use context to Terminate Goroutines (?)

Using the context to terminate a goroutine is a very common pattern.
It allows you to stop goroutines based on something from an earlier function in the call stack.

_example3_

### Turn Off a case in a select

Use `nil`

```go
package main

import "fmt"

func main() {
	done := make(chan struct{})
	in, in2 := make(chan int), make(chan int)

	go count(in, 10)
	go count(in2, 3)

	go func() {
		for {
			select {
			case v, ok := <-in:
				if !ok {
					in = nil // ignores on a next iteration
				} else {
					fmt.Println("in", v)
				}
			case v, ok := <-in2:
				if !ok {
					in2 = nil // ignores on a next iteration
				} else {
					fmt.Println("in2", v)
				}
			}

			if in == nil && in2 == nil {
				done <- struct{}{}
				return
			}
		}
	}()

	<-done // Wait for the completion signal
}

func count(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}
```

### Don't forget timeouts

Most interactive programs have to return a response within a certain amount of time.
One of the things that you can do with concurrency in Go is manage how much time a request
(or a part of a request) has to run. For example:

```go
func timeLimit[T any](worker func () T, limit time.Duration) (T, error) {
out := make(chan T, 1)
ctx, cancel := context.WithTimeout(context.Background(), limit)
defer cancel()
go func () {
out <- worker()
}()
select {
case result := <-out:
return result, nil
case <-ctx.Done():
var zero T
return zero, errors.New("work timed out")
}
}
```

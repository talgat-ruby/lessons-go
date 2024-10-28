# Lesson 12: Concurrency (WaitGroups and Mutex)

Sometimes one goroutine needs to wait for multiple goroutines to complete their work.
If you are waiting for a single goroutine, you can use the context cancellation pattern that you saw earlier.
But if you are waiting on several goroutines, you need to use a `WaitGroup`,
which is found in the sync package in the standard library.

```go
func main() {
var wg sync.WaitGroup
wg.Add(3)
go func () {
defer wg.Done()
doThing1()
}()
go func () {
defer wg.Done()
doThing2()
}()
go func () {
defer wg.Done()
doThing3()
}()
wg.Wait()
}
```

A `sync.WaitGroup` doesn’t need to be initialized, just declared, as its zero value is useful.
There are three methods on `sync.WaitGroup`:

- `Add`: which increments the counter of goroutines to wait for;
- `Done`: which decrements the counter and is called by a goroutine when it is finished;
- `Wait`: which pauses its goroutine until the counter hits zero.

`Add` is usually called once, with the number of goroutines that will be launched.
`Done` is called within the goroutine.
To ensure that it is called, even if the goroutine panics, you use a `defer`
It is important to pass the pointer of `WaitGroup`. If the pointer is not passed,
then each Goroutine will have its own copy of the `WaitGroup` and
main will not be notified when they finish executing.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
```

## Mutex

If you’ve had to coordinate access to data across threads in other programming languages,
you have probably used a `mutex`. This is short for mutual exclusion,
and the job of a mutex is to limit the concurrent execution of some code or access to a shared piece of data.
This protected part is called the _critical section_.

```go
x = x + 1
```

Internally the above line of code will be executed by the system
in the following steps(there are more technical details involving registers,
how addition works, and so on but for the sake of this tutorial lets assume that these are the three steps),

1. get the current value of x
2. compute x + 1
3. assign the computed value in step 2 to x

When these three steps are carried out by only one Goroutine, all is well.

```go
mutex.Lock()
x = x + 1
mutex.Unlock()
```

If one Goroutine already holds the _lock_ and if a new Goroutine is trying to acquire a _lock_,
the new Goroutine will be blocked until the `mutex` is unlocked.

### race condition

A _race condition_ occurs when two or more operations must execute in the correct order,
but the program has not been written so that this order is guaranteed to be maintained.

Most of the time, this shows up in what’s called a _data race_,
where one concurrent operation attempts to read a variable
while at some undetermined time another concurrent operation is attempting to write to the same variable.

_example1_

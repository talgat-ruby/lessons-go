# Lesson 13: Concurrency (RWMutex, sync, atomic)

> Do not communicate by sharing memory; instead, share memory by communicating

## RWMutex

`sync.RWMutex` helps synchronize write operations without synchronizing read operations,
while ensuring that there are no active read operations when there is a write operation in progress.

- Attempting to lock the **write** lock again when the **write** lock
  has been locked will **block** the current goroutine
- Attempting to lock the **read** lock again when the **write** lock is locked will also **block** the current goroutine
- Attempting to lock the **write** lock while the **read** lock is locked will also **block** the current goroutine
- Attempting to lock the **read** lock after the **read** lock has been locked will **not block** the current goroutine

_example1_

- If you are coordinating goroutines or tracking a value as it is transformed by a series of goroutines, use channels.
- If you are sharing access to a field in a struct, use mutexes.
- If you discover a critical performance issue when using channels,
  and you cannot find any other way to fix the issue, modify your code to use a mutex.

## sync.Map

When looking through the `sync` package, you’ll find a type called `Map`.
It provides a concurrency-safe version of Go’s built-in `map`. Because of trade-offs in its implementation,
`sync.Map` is appropriate only in very specific situations:

- When you have a shared map where key-value pairs are inserted once and read many times
- When goroutines share the map, but don’t access each other’s keys and values

Furthermore, because `sync.Map` was added to the standard library before the introduction of generics,
`sync.Map` uses any as the type for its keys and values.
This means the compiler cannot help you ensure that the right data types are used.
Given these limitations, in the rare situations where you need to share a `map` across multiple goroutines,
use a built-in `map` protected by a `sync.RWMutex`.

[More](https://sreramk.medium.com/go-inside-sync-map-how-does-sync-map-work-internally-97e87b8e6bf)

## sync.Once

It is used to perform a certain operation exactly once, regardless of how many times the `Once` instance is consulted.

```go
var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
once.Do(func () {
instance = createSingletonInstance()
})
return instance
}
```

Additionally, [more on `OnceFunc`, `OnceValue`, and
`OnceValues`.](https://reliasoftware.com/blog/new-features-in-golang-sync-once)

## sync.Condition

> Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an
> event.

`sync.Cond` could useful in situations where multiple readers wait for the shared resources to be available.

```go
var sharedRsc = make(map[string]interface{})
func main() {
var wg sync.WaitGroup
wg.Add(2)
m := sync.Mutex{}
c := sync.NewCond(&m)
go func () {
defer wg.Done()
// this go routine wait for changes to the sharedRsc
c.L.Lock()
defer c.L.Unlock()
for len(sharedRsc) == 0 {
c.Wait()
}
fmt.Println(sharedRsc["rsc1"])
}()

go func () {
defer wg.Done()
// this go routine wait for changes to the sharedRsc
c.L.Lock()
defer c.L.Unlock()
for len(sharedRsc) == 0 {
c.Wait()
}
fmt.Println(sharedRsc["rsc2"])
}()

// this one writes changes to sharedRsc
c.L.Lock()
sharedRsc["rsc1"] = "foo"
sharedRsc["rsc2"] = "bar"
c.Broadcast()
c.L.Unlock()
wg.Wait()
}
```

[More](https://kaviraj.me/understanding-condition-variable-in-go/)

## sync.Pool

A `sync.Pool` is a dynamic cache for objects,
primarily designed to optimize memory usage in concurrent programming scenarios in Go.
It allows you to reuse objects, reducing the overhead of frequent allocations and the strain on the garbage collector.
Goroutines can benefit from `sync.Pool` by accessing pooled objects without the need for continuous reallocation,
which is particularly advantageous in high-concurrency applications.

The utility of `sync.Pool` extends beyond scenarios with resource-intensive object creation,
like dealing with large data structures or managing shared resources like database connections;
it is also effective as a cache for already allocated resources, even for small and short-lived objects.
By reusing objects, `sync.Pool` helps optimize memory usage and
improve performance in applications that require frequent allocation and deallocation of objects.

### Pros and cons

**Pros:**

1. **Improved performance**: `sync.Pool` can significantly reduce the overhead of object creation and
   garbage collection, improving performance in scenarios where objects are frequently created and discarded.
2. **Resource efficiency**: It helps manage resources like memory,
   mainly when working with goroutines in high-concurrency environments,
   by reusing objects instead of continuously allocating new ones.
3. **Synchronization**: The pool handles synchronization internally,
   ensuring safe access to objects in concurrent environments, thus preventing race conditions.
   For example, when a goroutine retrieves or returns an object,
   **sync.Pool** uses these mechanisms to ensure that each operation is atomic,
   meaning it's completed fully without interruption.

**Cons:**

1. **Non-Persistence of Objects**: Objects in **sync.Pool** can be removed automatically during garbage collection
   without notification; this means you cannot rely on the persistence of objects in the pool over time.
2. **Overhead in low-concurrency scenarios**: In applications with low concurrency or where objects
   are not frequently reused, the overhead of managing a `sync.Pool` may not justify its benefits.
3. **Not suitable for different Object types**: The `sync.Pool` efficiency diminishes
   when handling a mix of different object types or sizes,
   as the pool cannot tailor its caching strategy to each object type's allocation and usage patterns,
   potentially leading to suboptimal memory usage and reduced performance gains.
4. **Potential code complexity**: Using `sync.Pool` could potentially increase the complexity of your codebase,
   making debugging harder due to unpredictable object reuse and
   non-deterministic garbage collection in concurrent scenarios.

```go
import "sync"

var workerPool = sync.Pool{}

// Initialize the pool with worker functions:
func init() {
workerPool.New = func () interface{} {
return func () {
// Perform some work in the goroutine
}
}
}

func main() {
// To execute work in a goroutine from the pool:
worker := workerPool.Get().(func())
go worker()

// When done, return the worker function to the pool:
workerPool.Put(worker)
}
```

## atomic

Atomic operations are operations that are designed to be executed without interruption and without interference
from other concurrent operations. They are used to ensure that certain operations on shared variables
are performed atomically, meaning they are executed as a single, indivisible unit,
and they are not subject to interference or data races from other goroutines or threads.

```go

package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

var counter int64

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for i := 0; i < 1000; i++ {
            atomic.AddInt64(&counter, 1)
        }
        wg.Done()
    }()
    
    go func() {
        for i := 0; i < 1000; i++ {
            atomic.AddInt64(&counter, 1)
        }
        wg.Done()
    }()
    
    wg.Wait()
    fmt.Println("Final counter value:", counter)
}
```

## Atomic vs Mutex

1. **Simplicity vs Complexity**: _Atomic_ operations are suitable for simple operations like incrementing a counter.
   _Mutexes_ are for more complex operations or multiple operations that need to be protected together.
2. **Performance**: _Atomic_ operations are generally faster than mutexes for their intended use-cases
   because they don't involve locking and blocking. However, overusing atomics, especially in non-trivial scenarios,
   can make code hard to understand and maintain.
3. **Read-heavy operations**: If you have a situation where there are numerous reads and occasional writes,
   _RWMutex_ can be beneficial. It allows multiple goroutines to read the data simultaneously but ensures
   exclusive access for writes.

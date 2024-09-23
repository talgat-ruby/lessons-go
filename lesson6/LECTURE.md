# Lesson 6: Function

## Function

You’ve already seen functions being declared and used. Every Go program starts from a main function.

```go
package main

import (
	"fmt"
)

func main() {
	result := div(5, 2)
	fmt.Println(result)
}

func div(num, denom int) int {
	return num / denom
}
```

### Variadic functions

The variadic parameter **must** be the last (or only) parameter in the input parameter list.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(addTo(3))             // []
	fmt.Println(addTo(3, 2))          // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
```

### Multiple return

The first difference that you’ll see between Go and other languages is that Go allows for _multiple return values_.

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder) // 2 1
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}
```

### Ignore values

Go **does not allow unused variables**. If a function returns multiple values, but you don’t need to read one or more of
the values, assign the unused values to the name `_`.

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	result, _, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result) // 2
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}
```

### Named return values

In addition to letting you return more than one value from a function, Go allows you to specify _names_ for your
return values.

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder) // 2 1
}

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	result, remainder, err = 0, 0, nil
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result = num / denom
	remainder = num % denom
	return
}
```

But please **don't use named return values**. They're confusing, unnecessary and considered antipattern.

### Functions as a value

```go
package main

import (
	"fmt"
)

var (
	add = func(i, j int) int { return i + j }
	sub = func(i, j int) int { return i - j }
	mul = func(i, j int) int { return i * j }
	div = func(i, j int) int { return i / j }
)

func main() {
	x := add(2, 3)
	fmt.Println(x) // 5
	changeAdd()
	y := add(2, 3) // 8
	fmt.Println(y)
}

func changeAdd() {
	add = func(i, j int) int { return i + j + j }
}
```

Also functions can be passed as parameters and as a return value.

```go
package main

import (
	"fmt"
)

func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
```

### defer

The `defer` keyword is used to delay the execution of a function or method until the surrounding function returns.
This allows you to ensure that certain operations are performed before a function returns, regardless of whether
an error occurs or not.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("result:", deferExample())
}

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}
```

```shell
$ go run ./...
exiting: 30
second: 20
first: 10
result: 30
```

### Call

Go is a _call-by-value_ language. It means that when you supply a variable for a parameter to a function,
Go **always makes a copy of the value** of the variable.

```go
package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func main() {
	i := 2
	s := "Hello"
	p := person{}
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	modifySuccess(&i, &s, &p)
	fmt.Println(i, s, p)
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modifySuccess(pi *int, ps *string, pp *person) {
	*pi = *pi * 2
	*ps = "Goodbye"
	pp.name = "Bob"
}
```

```shell
$ go run ./...
2 Hello {0 }
4 Goodbye {0 Bob}
```

#### map and slice

But because `map` and `slice` are both implemented as a **pointer to a struct**. They will be updated. Just remember
`map`
is always _mutable_, whereas `slice` works a _peculiar way_. Please read
[this blog post](https://www.pixelstech.net/article/1607859246-The-hidden-risk-of-passing-slice-as-function-parameter)
for more details.

```go
package main

import (
	"fmt"
)

func main() {
	m := map[int]string{
		1: "hello",
		2: "goodbye",
	}
	modifyMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println(s)
}

func modifyMap(m map[int]string) {
	m[2] = "hola"
	m[3] = "adios"
	delete(m, 1)
}

func modifySlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
```

```shell
$ go run ./...
map[2:hola 3:adios]
[2 4 6]
```

## Error

Basic operations:

- Design

```go
type error interface {
Error() string
}
```

- Construct

```go
errors.New("message")
```

```go
fmt.Errorf("message")
```

- Wrapping

```go
fmt.Errorf("message: %w", err)
```

```go
func (e CustomErr) Unwrap() error {
return e.Err
}
```

```go
// wrapping multiple errors

errors.Join(errs...)
```

- Check
  `Is` - If wrapped error matches **value**

```go
errors.Is(err, os.ErrNotExist)
```

`As` - If wrapped error matches **type**

```go
errors.As(err, &customErr{})
```

- Stack trace

```go
fmt.Printf("%+v", err)
```

### Panic and recover

**Panic** - it is a state generated by the Go runtime whenever it is unable to figure out what should happen next.
For example: read past the end of a `slice` or passing a negative size to `make`.

As soon as a panic happens, the current function exits immediately, and any defers attached to the current function
start running. When those defers complete, the defers attached to the calling function run, and so on, until main is
reached. The program then exits with a message and a stack trace.

You can create your own panic:

```go
package main

import (
	"os"
)

func main() {
	doPanic(os.Args[0])
}

func doPanic(msg string) {
	panic(msg)
}
```

Go provides a way to capture a panic to provide a more graceful shutdown or to prevent shutdown at all. The built-in
`recover` function is called from within a defer to check whether a panic happened. If there was a panic, the value
assigned to the `panic` is returned. Once a `recover` happens, execution continues normally.

```go
package main

import (
	"fmt"
)

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
```

```shell
$ go run ./...
60
30
runtime error: integer divide by zero
10
```

While `panic` and `recover` look a lot like exception handling in other languages,
they are not intended to be used that way. Reserve panics for fatal situations and use recover as a way to gracefully
handle these situations. If your program panics, be careful about trying to continue executing after the panic.
You’ll rarely want to keep your program running after a panic occurs.

## Generics

Generic types (simply called Generics) are codes that allow us to use them for various functions by just altering the
function types. Generics were created to make code independent of types and functions.

```go
package print

import (
	"fmt"
)

// for strings
func String(s []string) {
	for _, v := range s {
		fmt.Print(v)
	}
}

// for int
func Int(s []int) {
	for _, v := range s {
		fmt.Print(v)
	}
}
```

With Generics it would be, now it can accept any values

```go
package print

import (
	"fmt"
)

func Value[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}
```

### Constraints

in Go you have to define constraints of generic types. For example in stack implementation we are using `any`.

```go
package stack

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}
```

However, if we will extend with `Contains` method:

```go
package stack

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}
```

```shell
$ go run ./...
invalid operation: v == val (incomparable types in type set)
```

Because we can send any types. For example slices, which are not comparable with `==`. So we have to limit our types.

```go
type Stack[T comparable] struct {
vals []T
}
```

Custom constraints also can be created. However, they must be interfaced, which will be introduced later in our course.

### | and ~

Generics types can be defined as a set of multiple types with `|`.

```go
package main

import (
	"fmt"
)

func main() {
	sum1 := 104
	sum2 := 222
	minSum := Min(sum1, sum2)
	fmt.Println(minSum) // 104
}

func Min[T uint | int | int64](a, b T) T {
	if a < b {
		return a
	}
	return b
}
```

However, if we will define specific type and use it. It will cause errors.

```go
package main

import (
	"fmt"
)

type num uint

func main() {
	var sum1 num = 104
	var sum2 num = 222
	minSum := Min(
		sum1,
		sum2
	)                   // error num does not satisfy uint | int | int64 (possibly missing ~ for uint in uint | int | int64)
	fmt.Println(minSum) // 104
}

func Min[T uint | int | int64](a, b T) T {
	if a < b {
		return a
	}
	return b
}
```

To fix error we have to add `~`. Which basically says, use any underline type of current type.

```go
func Min[T ~uint | ~int | ~int64](a, b T) T {
if a < b {
return a
}
return b
}
```

It is so common we can use with `constraints`.

```go
func Min[T constraints.Integer](a, b T) T {
if a < b {
return a
}
return b
}
```

### Limitations

Although generics in Go have brought many benefits and new possibilities to the language, there are still some
limitations and challenges that come with their implementation. Here are some of the main limitations of generics in Go:

* Performance: One of the main concerns with generics in Go is the potential impact on performance. With the
  introduction of generics, the Go compiler needs to generate code for different types at compile time, which can lead
  to larger binaries and slower compilation times.
* Type constraints: Go's implementation of generics relies on type constraints to ensure type safety. However, these
  constraints can be restrictive and limit the types that can be used with generic functions and data structures.
* Syntax complexity: The syntax for declaring and using generic functions and data structures can be complex and
  difficult to understand, especially for beginners.
* Error messages: The error messages generated by the Go compiler for issues related to generics can be difficult to
  understand, making debugging and troubleshooting more challenging.
* Code readability: Generics in Go can sometimes make code less readable and harder to understand, especially if type
  constraints and type parameters are used extensively.
* No switching possible: When you want to switch from one underlying generic type to another, it is not possible using
  generics. The only way to go about this is to use an interface and run the type switch function at runtime.

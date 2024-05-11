# Lesson 1: Introduction

We will start learning basic structure of the go and we will write basic programs.

## A brief history

Go is a statically typed, compiled high-level programming language designed at Google 
by [Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer), [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike)
, and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson). It is syntactically similar to C, 
but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.
It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

## Prepare

```shell
$ go mod init github.com/talgat-ruby/lessons-go/lesson1

$ touch main.go
```

## Hello World

```go
// main.go

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

```shell
$ go run main.go
Hello, World!
```

## Greetings

```go
// main.go

package main

import "fmt"

func main() {
	greet("Bender", 5)
}

func greet(name string, table int) {
	text := fmt.Sprintf("Hello, %s! Your table is %d.\n", name, table)
	fmt.Println(text)
}
```

```shell
$ go run main.go
Hello, Bender! Your table is 5.
```

## Greet package

```go
// greet/greet.go

package greet

import (
	"fmt"
)

func Greet(name string, table int) string {
	return fmt.Sprintf("Hello, %s! Your table is %d.", name, table)
}
```

```go
// main.go

package main

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson1/greet"
)

func main() {
	text := greet.Greet("Leila", 5)

	fmt.Println(text)
}
```

```shell
$ go run main.go
Hello, Leila! Your table is 5.
```

## Build

```shell
$ go build

$ ./lesson1
Hello, Leila! Your table is 5.
```

```shell
$ go build -o greetings

$ ./greetings
Hello, Leila! Your table is 5.
```

## Debugging and Testing

For debugging, you can use [standart toolchain](https://go.dev/doc/gdb) or [Delve](https://github.com/go-delve/delve). But
I would suggest IDE (Goland, VS Code, etc.).

```go
// greet/greet_test.go

package greet

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, Leila! Your table is 5."
	text := Greet("Leila", 5)
	if text != expected {
		t.Errorf("Greetings was incorrect, got: %s, expected: %s.", text, expected)
	}
}

func TestTableGreet(t *testing.T) {
	table := []struct {
		name     string
		table    int
		expected string
	}{
		{"Leila", 5, "Hello, Leila! Your table is 5."},
		{"Zoidberg", -1, "Hello, Zoidberg! Your table is -1."},
		{"Prf. Farnsworth", 10, "Hello, Prf. Farnsworth! Your table is 10."},
	}

	for _, row := range table {
		text := Greet(row.name, row.table)
		if text != row.expected {
			t.Errorf("Greet(%s, %d) was incorrect, got: %s, expected: %s.", row.name, row.table, text, row.expected)
		}
	}
}
```

```shell
$ go test -v ./...
?   	github.com/talgat-ruby/lessons-go/lesson1	[no test files]
=== RUN   TestGreet
--- PASS: TestGreet (0.00s)
=== RUN   TestTableGreet
--- PASS: TestTableGreet (0.00s)
PASS
ok  	github.com/talgat-ruby/lessons-go/lesson1/greet	0.923s
```

## Homework:

Continue with previous task and create cli app, so it will accept name and table from user's terminal.

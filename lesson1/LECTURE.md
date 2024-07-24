# Lesson 1: Intro

<!-- TOC -->
* [Lesson 1: Intro](#lesson-1-intro)
  * [History](#history)
  * [Pros and cons](#pros-and-cons)
    * [Pros](#pros)
    * [Cons](#cons)
  * [Practice](#practice)
    * [Hello World](#hello-world)
  * [Exercises](#exercises)
    * [Exercise 1](#exercise-1)
    * [Exercise 2](#exercise-2)
<!-- TOC -->

## History

![history](./why_use_golang.png)

## Pros and cons

### Pros

* Fast
* Easy To Learn
* Well-Scaled
* Comprehensive

### Cons

* Time Consuming
* Less Features
* Binaries can be large

## Practice

```shell
$ go version
go version go1.22.5 darwin/arm64
```

### Hello World

```shell
$ mkdir example1 && cd example1

$ go mod init github.com/talgat-ruby/lessons-go/lesson1/example1

$ touch main.go
```

```go
// main.go

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Let's run the program

```shell
$ go run ./...
Hello, World!
```

Let's build the program and run binary file

```shell
$ go build ./... -o hello_world
$ ./hello_world
```

## Exercises

### Exercise 1

Build the `hello_world`, but now it will print your name instead of _World_

```shell
$ go run ./...
Hello, Talgat!
```

### Exercise 2

Use the code from **exercise 1**. And print number of time supplied from the cli argument. Default is **1**.

```shell
$ go run ./... 3
Hello, Talgat!
Hello, Talgat!
Hello, Talgat!
```

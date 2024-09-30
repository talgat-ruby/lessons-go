# Lesson 8: Extra

## Repo, Module & Packages

- Repository - a place in a version control system where the source code for a project is stored. _Modules_ are stored
  there.
- Module - a bundle of Go source code that’s distributed and versioned as a single unit. Can consist of several
  _packages_. Every module has a globally unique identifier.
- Package - a directory of a source code.

### First module

```shell
$ mkdir example1 && cd example1
$ go mod init github.com/talgat-ruby/lessons-go/lesson7/example1
```

Which will create `go.mod` file. Which is basically small description of your module.

### require

Directive `require` describes the dependencies of your project. Because we have none right now, we have to install in
order to see them:

```shell
$ go get github.com/rs/zerolog
```

```
// go.mod
module github.com/talgat-ruby/lessons-go /lesson7/example1

go 1.23.1

require (
  github.com/mattn/go -colorable v0.1.13 // indirect
  github.com/mattn/go -isatty v0.0.20 // indirect
  github.com/rs/zerolog v1.33.0 // indirect
  golang.org/x/sys v0.25.0 // indirect
)

```

- First `require` - direct dependencies which are used in the module
- Second `require` - indirect dependencies, either dependencies of dependencies or unused. Has a comment `//indirect`

Right now all our dependencies are indirect. Because we do not use them. Let's try to change it and create `main.go`

```shell
$ touch main.go
```

```go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")
}
```

Run it

```shell
$ go run main.go
```

If you check the `go.mod` nothing changed. Because we need to explain to "rewire" the description.

```shell
$ go mod tidy
```

```
// go.mod
module github.com/talgat-ruby/lessons-go /lesson7/example1

go 1.23.1

require github.com/rs/zerolog v1.33.0

require (
  github.com/mattn/go -colorable v0.1.13 // indirect
  github.com/mattn/go -isatty v0.0.20 // indirect
  golang.org/x/sys v0.25.0 // indirect
)
```

But if we will stop using it:

```go
// main.go
package main

import (
	"context"
	"log/slog"
)

func main() {
	ctx := context.Background()
	slog.InfoContext(ctx, "hello world")
}
```

```shell
$ go mod tidy
```

```
// go.mod
module github.com/talgat-ruby/lessons-go/lesson7/example1

go 1.23.1
```

### Importing & exporting

Let's create a package `math`. Inside `lesson8/example1`.

```shell
$ mkdir math && cd math
$ touch main.go
```

```go
// math/main.go
package math

func Add(a, b int) int {
	return a + b
}
```

Now we can use it in `main`.

```go
// main.go
package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/lesson7/example1/math"
)

func main() {
	ctx := context.Background()
	slog.InfoContext(ctx, fmt.Sprintf("%d + %d is %d", 3, 5, math.Add(3, 5)))
}
```

### internal package

The exported identifiers in that package and its subpackages are accessible only to the direct parent package of
internal and the sibling packages of `internal`.

### Circular Dependencies

Go does not allow you to have a circular dependency between packages. If `package A` imports `package B`, directly or
indirectly, `package B` cannot import `package A`, directly or indirectly.

If you find yourself with a circular dependency, you have a few options. In some cases, this is caused by splitting up
packages too finely. If two packages depend on each other, there’s a good chance they should be merged into a single
package. You can merge the `A` and `B` packages into a single package, and that solves your problem.

If you have a good reason to keep your packages separated, it may be possible to move just the items that cause the
circular dependency to one of the two packages or to a new package.

## Config

### Makefile

Modern software development relies on repeatable, automatable builds that can be run by anyone, anywhere, at any time.
The way to do this is to use some kind of script to specify your build steps.
Go developers have adopted `make` as their solution.
It lets developers specify a set of operations that are necessary to build a program and the order in which the steps
must be performed. You may not be familiar with `make`, but it’s been used to build programs on Unix systems since 1976.

```makefile
.DEFAULT_GOAL := user

.PHONY:user
user:
	echo $(USER)
```

```shell
$ make
echo talgatsaribayev
talgatsaribayev
```

For more info please use the [Makefile tutorial](https://makefiletutorial.com/#getting-started).

However, there are 2 issues with Makefiles:

- Learning curve for more complicated commands
- Is not natively supported on Windows machines

### Secondary tools

There are several alternatives to `make`

1. [Cmake](https://cmake.org/)
2. [Taskfile](https://taskfile.dev/) *
3. [just](https://just.systems/)
4. Or any other tool that helps with cli.

### cli and environment variables

**!Proceed with examples.**

## Go structure

### `main` package

`main` package is default one run. Use `main.go` file to keep as an entry point.

### internal

When you create a package called `internal`, the exported identifiers in that package and its subpackages are
accessible only to the direct parent package of internal and the sibling packages of internal.

### cmd

For library modules to have one or more applications included with them as utilities.

### bin

For executables.

### pkg

For packages which could be used by other packages. Try not to use it.

## Init function

Avoid it! Don't use it!

When you declare a function named `init` that takes no parameters and returns no values,
it runs the first time the package is referenced by another package. Since init functions do not have any
inputs or outputs, they can work only by side effect, interacting with package-level functions and variables.

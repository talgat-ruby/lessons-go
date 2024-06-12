# Lesson 7: Additional Go

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

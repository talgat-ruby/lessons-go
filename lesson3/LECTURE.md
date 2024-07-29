# Lesson 3: Basic operations

## Block

Each place where a declaration occurs is called a block. Variables, constants, types, 
and functions declared outside any functions are placed in the package block. Within a function, 
every set of braces ({}) defines another block, 
and in a bit you will see that the control structures in Go define blocks of their own.

You can access an identifier defined in any outer block from within any inner block. 
This raises the question: 
what happens when you have a declaration with the same name as an identifier in a containing block? 
If you do that, you _shadow_ the identifier created in the outer block.

## Shadowing Variables

```go
func main() {
    x := 10
    if x > 5 {
        fmt.Println(x) // 10
        x := 5
        fmt.Println(x) // 5
    }
    fmt.Println(x) // 10
}
```

**!example1**

A _shadowing variable_ is a variable that has the same name as a variable in a containing block. 
For as long as the shadowing variable exists, you cannot access a shadowed variable.

You can also shadow the package or default functions:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt := "oops"
	fmt.Println("hello")
}
```

```shell
$ go run ./...
./main.go:10:6: fmt.Println undefined (type string has no field or method Println)
```

## If, else if, else

The `if` statement in Go is much like the `if` statement in most programming languages.

```go
n := rand.Intn(10) 
if n < 5 {
	fmt.Println("That's too low:", n) 
} else if n > 5 {
	fmt.Println("That's too big:", n) 
} else {
    fmt.Println("Perfect:", n)
}
```

Go adds is the ability to declare variables that are scoped to the condition and to both the `if` and `else` blocks.

```go
if n := rand.Intn(10); n < 5 {
	fmt.Println("That's too low:", n)
} else if n > 5 {
	fmt.Println("That's too big:", n)
} else {
	fmt.Println("Perfect:", n)
}
```

## Loop

* C-style for with open parts
```go
for i := 0; i < 10; i++ { 
	fmt.Println(i)
}
```
```go
i := 0
for ; i < 10; i++ {
    fmt.Println(i)
}
```
```go
for i := 0; i < 10; { 
	fmt.Println(i)
	if i % 2 == 0 { 
		i++
	} else { 
		i+=2
	}
}
```
```go
i := 1
for i < 100 {
	fmt.Println(i)
	i=i*2 
}
```
```go
for {
	fmt.Println("Hello")
}
```
* for-range

```go
fibNums := []int{1, 1, 2, 3, 5, 8} 
for i, v := range fibNums {
    fmt.Println(i, v)
}
```

**!example2**

### break and continue

`break` exits the loop immediately, just like the `break` statement in other languages

```go
i := 0
for {
    if i == 5 {
        break
    }
	i++
}
fmt.Println(i)
```

`continue` keyword, which skips over the rest of the for loop’s body and proceeds directly to the next iteration

```go
for i := 1; i <= 100; i++ {
    if i%3 == 0 {
        if i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else {
            fmt.Println("Fizz")
        }
    } else if i%5 == 0 {
        fmt.Println("Buzz")
    } else {
        fmt.Println(i)
    }
}
```

But it is not idiomatic

```go
for i := 1; i <= 100; i++ { 
    if i%3 == 0 && i%5 == 0 {
        fmt.Println("FizzBuzz")
        continue
    }
    if i%3 == 0 {
        fmt.Println("Fizz")
        continue
    }
    if i%5 == 0 {
        fmt.Println("Buzz")
        continue
    }
    fmt.Println(i)
}
```

### for-range value is a copy

**!example3**

## Labels and goto

**!example4**

## switch

Most developers in those languages avoid `switch` statements because of their limitations on values 
that can be switched on and the default fall-through behavior. But Go is different.
It makes `switch` statements useful.

```go
words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
for _, word := range words {
    switch size := len(word) {
	case size < 5:
        fmt.Println(word, "is a short word!")
    case 5:
        wordLen := len(word)
        fmt.Println(word, "is exactly the right length:", wordLen)
    case size > 10:
    default:
        fmt.Println(word, "is a long word!")
    }
}
```

```go
words := []string{"hi", "salutations", "hello"} 
for _, word := range words {
    switch wordLen := len(word); { 
	case wordLen < 5:
        fmt.Println(word, "is a short word!")
    case wordLen > 10:
		fmt.Println(word, "is a long word!")
    default:
        fmt.Println(word, "is exactly the right length.")
    }
}
```

!fallthrough

## Exercises

### Exercise 1

Write a for loop that puts 100 random numbers between 0 and 100 into a slice and prints it.

### Exercise 2

Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
1. If the value is divisible by 2, print “Two!” 
2. If the value is divisible by 3, print “Three!” 
3. If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else. 
4. Otherwise, print “Never mind”.

## Exercise 3

Start a new program. In main, declare an `int` variable called `total`. 
Write a `for` loop that uses a variable named `i` to iterate from `0` (inclusive) to `10` (exclusive). 
The body of the for loop should be as follows:
```go
total := total + i
fmt.Println(total)
```
After the for loop, print out the value of `total`.

# Lesson 2: Primitive types

## var vs :=

The most verbose way to declare a variable in Go uses the `var` keyword, an explicit type, and an assignment. 
It looks like this:

```go
var x int = 10
```

If the type on the righthand side of the = is the expected type of your variable, 
you can leave off the type from the left side of the =. Since the default type of an integer literal is int, 
the following declares x to be a variable of type `int`:

```go
var x = 10
```

If you are declaring multiple variables at once, you can wrap them in a declaration list:

```go
var (
    x int
    y = 20
    z int = 30
    d, e = 40, "hello" 
    f, g string
)
```

When you are **within a function**, 
you can use the `:=` operator to replace a `var` declaration that uses type inference. 
The following two statements do exactly the same thing—they declare `x` to be an `int` with the value of `10`:

```go
var x = 10 
x := 10
```

## const

Many languages have a way to declare a value as immutable. In Go, this is done with the `const` keyword.

```go
package main

import "fmt"

const x int64 = 10

const (
    idKey = "id"
    nameKey = "name"
)

const z = 20 * 10 

func main() {
    const y = "hello" 
    fmt.Println(x)
    fmt.Println(y)
    x = x + 1 // this will not compile!
    y = "bye" // this will not compile!
    fmt.Println(x)
    fmt.Println(y)
}
```

* Numeric literals 
* true and false 
* Strings 
* Runes 
* The values returned by the built-in functions complex, real, imag, len, and cap 
* Expressions that consist of operators and the preceding values


## Boolean

```go
var flag bool // no value assigned, set to false 
var isAwesome = true
```

## Numeric

### Integer

| Type name | Value range                                 |
|:----------|---------------------------------------------|
| `int8`    | –128 to 127                                 |
| `int16`   | –32768 to 32767                             |
| `int32`   | –2147483648 to 2147483647                   |
| `int64`   | –9223372036854775808 to 9223372036854775807 |
| `uint8`   | 0 to 255                                    |
| `uint16`  | 0 to 65535                                  |
| `uint32`  | 0 to 4294967295                             |
| `uint64`  | 0 to 18446744073709551615                   |

#### Special integers

| Alias  | Value                |
|--------|----------------------|
| `byte` | `uint8`              |
| `int`  | `int32` or `int64`   |
| `uint` | `uint32` or `uint64` |
| `rune` | `int32`              |

### Float

| Type name | Largest absolute value                                  | Smallest (nonzero) absolute value              |
|:----------|---------------------------------------------------------|------------------------------------------------|
| `float32` | 3.40282346638528859811704183484516925440e+38            | 1.401298464324817070923729583289916131280e-45  |
| `float64` | 1.797693134862315708145274237317043567981e+308          | 4.940656458412465441765687928682213723651e-324 |

### Complex

`complex64` uses `float32` values to represent the real and imaginary part, and `complex128` uses `float64` values.

```go
var x = complex(2.5, 3.1)
```

### Type Conversion

```go
package main

import "fmt"

func main() {
	var x = 10
	var y = 30.2
	var sum1 = float64(x) + y
	var sum2 = x + int(y)
	fmt.Println(sum1, sum2) // 40.2 40
}
```

## String

```go
var x string = "Hello"
fmt.Println(x)
```

## Exercises

### Exercise 1

Write a program that declares an integer variable called `i` with the value `20`. 
Assign `i` to a floating-point variable named `f`. Print out `i` and `f`.

### Exercise 2

Write a program with three variables, one named `b` of type `byte`, one named `smallI` of type `int32`, 
and one named `bigI` of type `uint64`. 
Assign each variable the maximum legal value for its type; then add `1` to each variable. Print out their values.

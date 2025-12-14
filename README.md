# Gofuncy

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-see%20LICENSE-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/CZero/gofuncy)](https://goreportcard.com/report/github.com/CZero/gofuncy)

My library of functions I keep using / writing.

## Table of Contents

| Package | Description | Functions |
|---------|-------------|-----------|
| [Errors](#errors) | Error handling utilities | `PanErr` |
| [Strings](#strings) | String manipulation | `SilentAtoi`, `ReverseString`, `RemoveSpaces` |
| [Slices](#slices) | Slice operations | `SumSlice`, `ProductSlice` |
| [Maths](#maths) | Mathematical functions | |
| └ [Digits](#digits) | Digit range helpers | `FindMinMaxInxDigits` |
| └ [Division](#division) | Divisor calculations | `Getdivisors` |
| └ [Factorial](#factorial) | Factorial calculations | `GetFactorial` |
| └ [Triangle Numbers](#triangle-numbers) | Triangle number calculations | `Gettrianglenumber` |
| └ [Primes](#primes) | Prime number utilities | Internal helpers |
| [Matrix](#matrix) | String grid operations | Word search, matrix types |
| [File Functions](#file-functions) | File I/O utilities | `ReadLines` |

## Installation

```bash
go get github.com/CZero/gofuncy
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/CZero/gofuncy/stringsfuncs"
    "github.com/CZero/gofuncy/mathfuncs/factorialfuncs"
    "github.com/CZero/gofuncy/slicesfuncs"
)

func main() {
    // String operations
    reversed := stringsfuncs.ReverseString("hello")
    fmt.Println(reversed) // "olleh"

    // Math operations
    fact := factorialfuncs.GetFactorial(5)
    fmt.Println(fact) // 120

    // Slice operations
    numbers := []int{1, 2, 3, 4, 5}
    sum := slicesfuncs.SumSlice(numbers)
    fmt.Println(sum) // 15
}
```

## Packages

### Errors

```go
import "github.com/CZero/gofuncy/errorsfuncs"
```

#### PanErr

Panics on Error

```go
func PanErr(err error)
```

### Strings

```go
import "github.com/CZero/gofuncy/stringsfuncs"
```

#### SilentAtoi

Returns an int or panics on conversion failure

```go
func SilentAtoi(input string) int
```

#### ReverseString

Returns a reversed string

```go
func ReverseString(word string) string
```

#### RemoveSpaces

Removes all extra whitespace from a string

```go
func RemoveSpaces(input string) string
```

### Slices

```go
import "github.com/CZero/gofuncy/slicesfuncs"
```

#### SumSlice

Returns the sum of an integer slice

```go
func SumSlice(inOrder []int) int
```

#### ProductSlice

Returns the product of an integer slice

```go
func ProductSlice(inOrder []int) int
```

### Maths

Mathematical utility functions organized by topic.

#### Digits

```go
import "github.com/CZero/gofuncy/mathfuncs/digitsfuncs"
```

##### FindMinMaxInxDigits

Returns the smallest and largest numbers with the requested number of digits.

```go
func FindMinMaxInxDigits(digits int) (min, max int)
```

#### Division

```go
import "github.com/CZero/gofuncy/mathfuncs/divisionfuncs"
```

##### Getdivisors

Finds all divisors of a natural number efficiently using square root optimization

```go
func Getdivisors(i int) []int
```

#### Factorial

```go
import "github.com/CZero/gofuncy/mathfuncs/factorialfuncs"
```

##### GetFactorial

Returns the factorial of a number. Returns 0 for input ≤ 0

```go
func GetFactorial(input int) int
```

#### Triangle Numbers

```go
import "github.com/CZero/gofuncy/mathfuncs/trianglenumbersfuncs"
```

##### Gettrianglenumber

Gets the n'th triangle number (sum of natural numbers from 1 to n). For example, the 7th triangle number is 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28

```go
func Gettrianglenumber(n int) int
```

#### Primes

```go
import "github.com/CZero/gofuncy/mathfuncs/primesfuncs"
```

Prime number utilities currently exposing internal helpers for building new functionality.

##### getPrimesBelow *(internal helper)*

Calculates all prime numbers below a given limit.

```go
func getPrimesBelow(limit int) []int
```

##### getPrimeFactors *(internal helper)*

Finds the prime factors of a number by iteratively dividing by primes.

```go
func getPrimeFactors(number int) []int
```

> These helpers are not yet exported; wrap them in your own functions or contribute an exported API if you need direct access from another module.

### Matrix

```go
import "github.com/CZero/gofuncy/matrixfuncs"
```

Provides matrix operations for string grids, including word search functionality in all directions (horizontal, vertical, diagonal).

#### Types

```go
type Coord struct {
    r int // row
    c int // col
}

type Matrix map[Coord]string

type StringsMatrix struct {
    grid   Matrix
    height int
    width  int
}
```

### File Functions

```go
import "github.com/CZero/gofuncy/filefuncs"
```

#### ReadLines

Reads a whole file into memory and returns a slice of its lines

```go
func ReadLines(path string) ([]string, error)
```

## Requirements

- Go 1.25.1 or higher
- No external dependencies

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the terms found in the [LICENSE](LICENSE) file.

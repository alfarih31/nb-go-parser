# Noob-Parser

[![Go Reference](https://pkg.go.dev/badge/github.com/alfarih31/nb-go-parser.svg)](https://pkg.go.dev/github.com/alfarih31/nb-go-parser)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/alfarih31/nb-go-parser?style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/alfarih31/nb-go-parser?style=flat-square)

Noob Parser is a tools to parser Primitive types, such as:
- String masking
- Cast `String` to `Array of String`
- Cast `String` to `Array of Int`
- Cast `String` to `Bool` & Vice Versa
- Cast `String` to `Int` & Vice Versa
- Cast `Int` to `Bool` & Vice Versa
- Cast `[]int32` to `[]int` & Vice Versa 
- Cast `[]int64` to `[]int` & Vice Versa 

## Contents

- [Noob-Parser](#Noob-Parser)
  - [Installation](#Installation)
  - [Quick Start & Usage](#Quick)

## Installation

To install this package, you need to install Go (**version 1.17+ is required**) & initiate your Go workspace first.

1. After you initiate your workspace then you can install this package with below command.

```shell
go get -u github.com/alfarih31/nb-go-parser
```

2. Import it in your code

```go
import "github.com/alfarih31/nb-go-parser"
```

## Quick Start & Usage

See the test:
- [string_test.go](string_test.go)
- []

## Contributors ##

- Alfarih Faza <alfarihfz@gmail.com>

## License

This project is licensed under the - see the [LICENSE.md](LICENSE.md) file for details
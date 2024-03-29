# goastutil

The `goastutil` package provides utility functions for working with Go abstract
syntax tree (AST).


# Comparison with `Jennifier` Package

[`Jennifier`](https://github.com/dave/jennifer) is a good and awesome code generator tool for Go. In most of cases
you should use it rather than [`goastutil`](https://github.com/septemhill/goastutil). Unless you want to do some Source-to-Source (from Go to Go) code transformation and generation from AST, then we serve for you.

# Installation

To install the `goastutil` package, run the following command:

```shell
go get github.com/septemhill/goastutil
```

# Usage

To use the goastutil package, import it in your Go code:

```go
import "github.com/septemhill/goastutil"
```

You can then use the utility functions provided by the goastutil package in your
code.

# Examples

## Add / Remove Declarations

```go
// sample.go
package main

type Base struct {
	Name string
}

func (b *Base) PointerReceiver() {}

func (b Base) ConcreteReceiver() {}

func existFunc(a, b int) int {
	return a + b
}
```
```go
// main.go
package main

import (
  "log"
  "github.com/septemhill/goastutil"
)

func main() {
  astFile, err := goastutil.NewFile("./sample.go")
  if err != nil {
  	log.Fatal("error: ", err)
  }

  // Append `func fnDecl` into the astFile.
  astFile.AddDecls(`
    func fnDecl() {}
  `)

  // Append multiple decls: `const constDecl = 1` and `var varDecl = "varDecl"` into the astFile.
  astFile.AddDecls(`
  	const constDecl = 1
  	var varDecl = "varDecl"
  `)

  // Remove `existFunc`, `constDecl` (which just created above) and 
  // Base's methods from the astFile. But the `Base.PointerReceiver` 
  // wouldn't be removed, we should use `*Base.PointerReceiver`
  astFile.RemoveDecls([]string{
    "existFunc", 
    "constDecl", 
    "Base.PointerReceiver", 
    "Base.ConcreteReceiver",
  })

  // After operations above, astFile should be like:
  // package main
  //
  // var varDecl = "varDecl"
  //
  // type Base struct {
  // 		Name string
  // }
  //
  // func (b *Base) PointerReceiver() {
  // }
  //
  // func fnDecl() {
  // }
  println(astFile.String())
}

```

## Add / Remove Imports
```go
// sample.go
package main

import "fmt"

func printInt(x int) {
	fmt.Println(x)
}
```
```go
package main

import (
  "log"
  "github.com/septemhill/goastutil"
)

func main() {
  astFile, err := goastutil.NewFile("./sample.go")
  if err != nil {
  	log.Fatal("error: ", err)
  }

  // Append `errors` into the astFile
  astFile.AddImports([]string{
    "errors",
  })

  // Remove `fmt` from the astFile
  astFile.RemoveImports([]string{
    "fmt",
  })

  // After remove `fmt`, astFile should be like:
  // package main
  //
  // import "errors"
  //
  // func printInt(x int) {
  //         fmt.Println(x)
  // }
  println(astFile.String())
}
```
## Add / Remove Statements
```go
package main

func existFunc(a int, b int) int {
  return a + b
}
```
```go
package main

import (
  "log"
  "github.com/septemhill/goastutil"
)

func main() {
  astFile, err := goastutil.NewFile("./sample.go")
  if err != nil {
  	log.Fatal("error: ", err)
  }

  // Prepend a stmt to the first function body
  astFile.
    FuncDecls()[0].
    Body().
    Stmts()[0].
    PrependStmt(`
      fmt.Println(a, b)
    `)

  // Prepend multiple stmts to the first function 
  // body (which is `fmt.Println(a, b)`)
  astFile.
    FuncDecls()[0].
    Body().
    Stmts()[0].
    PrependStmt(`
  	  const prepend = "Septem"
  	  var append = 123456
    `)

  // Remove the first stmt in the first function body
  // (which is `const prepend = "Septem"`)
  astFile.
    FuncDecl()[0].
    Body().
    RemoveStmt(0, 1)

  // After operations above, astFile should be like:
  // package main
  //
  // func existFunc(a int, b int) int {
  // 		append := 123456
  // 		fmt.Println(a, b)
  // 		return a + b
  // }
  println(astFile.String())
}
```
# TODO

- [ ] Introduce AST modify functionality
  - [X] [Import](https://github.com/septemhill/goastutil/tree/main/example/imports)
    - [X] Add
    - [X] Remove
  - [X] [Declaration](https://github.com/septemhill/goastutil/tree/main/example/decls)
    - [X] Add
    - [X] Remove
  - [X] [Statement](https://github.com/septemhill/goastutil/tree/main/example/stmts)
    - [X] Prepend
    - [X] Append
    - [X] Remove
  - [ ] Expression
    - [ ] Modify value
    - [ ] Variable renaming
  - [ ] Comments ([Issue](https://github.com/golang/go/issues/20744))
  - [X] Struct Tags
- [X] Add generic support

# Contributing

Contributions to the `goastutil` package are welcome! If you find any issues or
have any improvements to suggest, please open an issue or submit a pull request.

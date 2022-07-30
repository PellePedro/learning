
# Reflection
[Reflection in Golang](https://go101.org/article/reflection.html)

#
https://stackoverflow.com/questions/64196547/is-possible-to-reflect-an-struct-from-ast

```
// main.go
package main

import (
    "bytes"
    "fmt"
    "go/ast"
    "go/parser"
    "go/printer"
    "go/token"
    "log"
)

// Main aa
func main() {
    // Create the AST by parsing src.
    fset := token.NewFileSet() // positions are relative to fset
    packages, err := parser.ParseDir(fset, "./entities", nil, 0)

    if err != nil {
        panic(err)
    }

    for _, pack := range packages {
        for _, file := range pack.Files {
            // Inspect the AST and print all identifiers and literals.
            ast.Inspect(file, func(n ast.Node) bool {
                switch x := n.(type) {
                case *ast.TypeSpec: // Gets Type assertions
                    fmt.Println(x.Name.Name)

                    v := x.Type.(*ast.StructType)
                    for _, field := range v.Fields.List {

                        for _, name := range field.Names {

                            // get field.Type as string
                            var typeNameBuf bytes.Buffer
                            err := printer.Fprint(&typeNameBuf, fset, field.Type)
                            if err != nil {
                                log.Fatalf("failed printing %s", err)
                            }
                            fmt.Printf("field %+v has type %+v\n", name.Name, typeNameBuf.String())
                        }

                    }
                    fmt.Println()

                }

                return true
            })
        }
    }
}
```


```
A long time passed and I find a way:
After you parsed a AST file and get the structs from package, you could use reflection to create a struct in runtime with the following:
t := reflect.StructOf([]reflect.StructField{
    {
        Name: "A",
        Type: reflect.TypeOf(int(0)),
        Tag:  `json:"a"`,
    },
    {
        Name: "B",
        Type: reflect.TypeOf(""),
        Tag:  `json:"B"`,
    },
    // Other fields ...
})

d := reflect.New(t).Interface() // Here you recived a struct as interface. And that's it.
```

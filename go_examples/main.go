package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) > 1 {
        name := os.Args[1]
        fmt.Printf("Hello, %s!\n", name)
    } else {
        fmt.Println("Hello, World!")
    }
}

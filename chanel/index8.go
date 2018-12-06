package main
import (
    "fmt"
)
func main() {
    defer func() {
        msg := recover()
        fmt.Println(msg)
    }()
    fmt.Println("I am walking and singing...")
    panic("It starts to rain cats and dogs")
}
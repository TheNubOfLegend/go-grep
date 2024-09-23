package main

import (
    "fmt"
    "os"
)

func main() {
    os.ReadFile("test")
    var pattern []byte = []byte("at that")
    var lookup []int = make([]int, 256)
    for i := len(pattern) - 1; i >= 0; i-- {
        if lookup[pattern[i]] == 0 {
            lookup[pattern[i]] = i
        }
    }
    for i := range lookup {
        if lookup[i] != 0 {
            fmt.Print(rune(i))
            fmt.Println(lookup[i])
        }
    }
}


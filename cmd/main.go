package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.ReadFile("test")
    if err != nil {
        fmt.Println("no file nerd")
    }

    var pattern []byte = []byte("at that")
    var lookup []int = make([]int, 256)
    for i := len(pattern) - 1; i >= 0; i-- {
        if lookup[pattern[i]] == 0 {
            lookup[pattern[i]] = len(pattern) - 1 - i
        }
    }
    for i := range lookup {
        if lookup[i] != 0 {
            fmt.Printf("%c", i)
            fmt.Println(lookup[i])
        }
    }
    offset := 0
    for i := len(pattern) - 1; i >= 0; i-- {
        delta := 0
        if file[i + offset] != pattern[i] {
            if lookup[file[i]] != 0 {
                delta += lookup[file[i + offset]]
            } else {
                delta += len(pattern)
            }
        } else {
            fmt.Println(offset)
            fmt.Printf("%c\n", file[i + offset])
        }
        if delta != 0 {
            offset += delta
            i = len(pattern) - 1
        }
    }
}

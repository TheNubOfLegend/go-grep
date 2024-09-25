package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.ReadFile("test2")
    if err != nil {
        fmt.Println("no file nerd")
    }
    // fmt.Printf("\n %d", len(file))

    var pattern []byte = []byte("at that")
    var lookup []int = make([]int, 256)
    for i := range lookup {
        lookup[i] = len(pattern)
    }
    for i := len(pattern) - 1; i >= 0; i-- {
        if lookup[pattern[i]] == len(pattern) {
            lookup[pattern[i]] = len(pattern) - 1 - i
        }
    }
    for i := range lookup {
        if lookup[i] != len(pattern) {
            fmt.Printf("%c", i)
            fmt.Println(lookup[i])
        }
    }
    offset := 0
    count := 0
    for offset < len(file) + len(pattern) {
        for i := len(pattern) - 1; i >= 0 && i + offset < len(file); i-- {
            delta := 0
            if file[i + offset] != pattern[i] {
                if lookup[file[i]] != len(pattern) {
                    delta += lookup[file[i + offset]]
                } else {
                    delta += len(pattern)
                }
            } else {
                if i == 0 {
                    delta += len(pattern)
                    for j := 0; j < len(pattern); j++ {
                        fmt.Printf("%c", file[offset + j])
                    }
                    count++
                    fmt.Printf("\n%d found pattern at %d\n", count, offset)
                    // fmt.Println()
                    // j := 0
                    // k := 0
                    // for j < 50 {
                    //     if file[offset + j] == '\n' {
                    //         break
                    //     }
                    //     j++
                    // }
                    // for k > -50 {
                    //     if file[offset + k] == '\n' {
                    //         break
                    //     }
                    //     k--
                    // }
                    // for l := offset + k; l < offset + j; l++ {
                    //     fmt.Printf("%c", file[offset + l])
                    // }
                }
                // fmt.Println(offset)
                // fmt.Printf("%c\n", file[i + offset])
            }
            if delta != 0 {
                offset += delta
                i = len(pattern) - 1
            }
            // fmt.Println(offset)
        }
    }
}

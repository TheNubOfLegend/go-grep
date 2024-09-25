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
    fmt.Printf("\n %d", len(file))

    var pattern []byte = []byte("at that")
    var lookup []int = make([]int, 256)
    for i := len(pattern) - 1; i >= 0; i-- {
        if lookup[pattern[i]] == 0 {
            //shift one less, made this way to avoid initing the entire array to max shift
            lookup[pattern[i]] = len(pattern) - i
        }
    }
    for i := range lookup {
        if lookup[i] != 0 {
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
                if lookup[file[i]] != 0 {
                    delta += lookup[file[i + offset]] - 1
                } else {
                    delta += len(pattern)
                }
            } else {
                if i == 0 {
                    delta += len(pattern)
                    for j := 0; j < len(pattern); j++ {
                        fmt.Printf("%c", file[offset + j])
                    }
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
                    count++
                    fmt.Printf("\n%d found pattern at %d\n", count, offset)
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

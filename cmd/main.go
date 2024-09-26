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
    // fmt.Printf("\n %d\n", len(file))

    var pattern []byte = []byte("at that")
    var lookup []int = make([]int, 256)
    // for i := range lookup {
    //     lookup[i] = len(pattern)
    // }
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
    // panic("nerd")
    // for i := 72540; i < 72580; i++ {
    //     fmt.Printf("%c", file[i])
    // }
    fmt.Println()
    offset := 0
    count := 0
    delta := 0
    i := len(pattern) - 1
    for offset + len(pattern) < 72580 {
        for i >= 0 && i + offset < 72580 {
            if delta != 0 {
                offset += delta + len(pattern) - 1 - i
                i = len(pattern) - 1
                delta = 0
            }
            fmt.Printf("%c ", file[i+offset])
            fmt.Println(i + offset)
            if file[i + offset] != pattern[i] {
                delta += lookup[file[i + offset]] - 1
            } else if i == 0 {
                delta += len(pattern)
                // for j := 0; j < len(pattern); j++ {
                //     fmt.Printf("%c", file[offset + j])
                // }
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
            fmt.Printf("    d %d, i %d\n", delta, i)
            // fmt.Println(offset)
            // fmt.Printf("%c\n", file[i + offset])
            i--
        }
        // fmt.Println(offset)
    }
}

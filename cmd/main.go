package main

import (
	"fmt"
	"os"
)

func main() {
    horspool()
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func horspool() {
    file, err := os.ReadFile("test")
    if err != nil {
        fmt.Println("no file nerd")
    }

    var pattern []byte = []byte("good")
    var badChar [256]int

    for i := range pattern {
        badChar[pattern[i]] = i
    }

    offset := 0
    count := 0
    i := len(pattern) - 1
    for offset + len(pattern) < len(file) {
        if file[i + offset] != pattern[i] {
            offset += max(1, i - badChar[file[i + offset]])
            i = len(pattern) - 1
        } else if i == 0 {
            offset += len(pattern)
            i = len(pattern) - 1

            count++
            fmt.Printf("%d found pattern at %d\n", count, offset)
        } else {
            i--
        }
    }
}

func BM() {
    file, err := os.ReadFile("test")
    if err != nil {
        fmt.Println("no file nerd")
    }

    var pattern []byte = []byte("good")
    var badChar [256]int

    for i := range pattern {
        badChar[pattern[i]] = i
    }

    offset := 0
    count := 0
    i := len(pattern) - 1
    for offset + len(pattern) < len(file) {
        if file[i + offset] != pattern[i] {
            offset += max(1, i - badChar[file[i + offset]])
            i = len(pattern) - 1
        } else if i == 0 {
            offset += len(pattern)
            i = len(pattern) - 1

            count++
            fmt.Printf("%d found pattern at %d\n", count, offset)
        } else {
            i--
        }
    }
}

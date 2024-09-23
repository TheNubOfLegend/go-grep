package main

import (
    "fmt"
    "os"
)

func main() {
    os.ReadFile(test)
    var pattern []byte
    var lookup [256]byte
    for i := len(pattern) - 1; i >= 0; i++ {
        if lookup[pattern[i]] == 0 {
            lookup[pattern[i]] = i
        }
    }


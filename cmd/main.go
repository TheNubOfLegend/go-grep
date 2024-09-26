package main

import (
	"fmt"
	"os"
)

//
// func main() {
//     file, err := os.ReadFile("test")
//     if err != nil {
//         fmt.Println("no file nerd")
//     }
//     // fmt.Printf("\n %d\n", len(file))
//
//     var pattern []byte = []byte("at that")
//     var badChar []int = make([]int, 256)
//     // for i := range badChar {
//     //     badChar[i] = len(pattern)
//     // }
//     for i := len(pattern) - 1; i >= 0; i-- {
//         if badChar[pattern[i]] == 0 {
//             badChar[pattern[i]] = len(pattern) - 1 - i
//         }
//     }
//     for i := range badChar {
//         if badChar[i] != 0 {
//             fmt.Printf("%c", i)
//             fmt.Println(badChar[i])
//         }
//     }
//     // panic("nerd")
//     // for i := 72540; i < 72580; i++ {
//     //     fmt.Printf("%c", file[i])
//     // }
//     fmt.Println()
//     offset := 0
//     count := 0
//     delta := 0
//     i := len(pattern) - 1
//     for offset + len(pattern) < 72580 {
//         for i >= 0 && i + offset < 72580 {
//             if delta != 0 {
//                 offset += delta //+ len(pattern) - 1 - i
//                 i = len(pattern) - 1
//                 delta = 0
//             }
//             fmt.Printf("%c ", file[i+offset])
//             fmt.Println(i + offset)
//             if file[i + offset] != pattern[i] {
//                 delta += badChar[file[i + offset]]
//             } else if i == 0 {
//                 delta += len(pattern)
//                 // for j := 0; j < len(pattern); j++ {
//                 //     fmt.Printf("%c", file[offset + j])
//                 // }
//                 count++
//                 fmt.Printf("\n%d found pattern at %d\n", count, offset)
//                 // fmt.Println()
//                 // j := 0
//                 // k := 0
//                 // for j < 50 {
//                 //     if file[offset + j] == '\n' {
//                 //         break
//                 //     }
//                 //     j++
//                 // }
//                 // for k > -50 {
//                 //     if file[offset + k] == '\n' {
//                 //         break
//                 //     }
//                 //     k--
//                 // }
//                 // for l := offset + k; l < offset + j; l++ {
//                 //     fmt.Printf("%c", file[offset + l])
//                 // }
//             }
//             fmt.Printf("    d %d, i %d\n", delta, i)
//             // fmt.Println(offset)
//             // fmt.Printf("%c\n", file[i + offset])
//             i--
//         }
//         // fmt.Println(offset)
//     }
// }

func main() {
    file, err := os.ReadFile("test")
    if err != nil {
        fmt.Println("no file nerd")
    }
    fmt.Println(len(file))

    var pattern []byte = []byte("at that")
    var badChar [256]int
    // var goodSuff []int = make([]int, len(pattern))

    for i := range badChar {
        badChar[i] = len(pattern)
    }

    for i := range pattern {
        badChar[pattern[i]] = len(pattern) - 1 - i
    }

    offset := 0
    count := 0
    i := len(pattern) - 1
    for offset + len(pattern) < 455080 {
        bump := false
        delta := 0
        if file[i + offset] != pattern[i] {
            delta = badChar[file[i + offset]]
            if delta == 0 {
                bump = true
            }
        } else if i == 0 {
            delta = len(pattern)
            count++
            fmt.Printf("%d found pattern at %d\n", count, offset)
        } else {
            i--
        }
        if delta != 0 {
            offset += delta
            i = len(pattern) - 1
        } else if bump {
            offset += 1
            // i = len(pattern) - 1
        }
        fmt.Printf("    d %d, i %d\n", delta, i)
        fmt.Println(offset)
        fmt.Printf("%c\n", file[i + offset])
    }
}

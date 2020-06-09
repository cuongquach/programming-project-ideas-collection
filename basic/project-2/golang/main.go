package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // Get list string input args
    args_input := os.Args[1:]

    // Concentrate array input args to string
    str1 := strings.Join(args_input, " ")

    // Split strings each words to array again
    list_each_words := strings.Split(str1, "")

    // List vowels: A, E, I, O, U
    var count_a, count_e, count_i, count_o, count_u int = 0, 0, 0, 0, 0

    for _, element := range list_each_words {

        element = strings.ToLower(element)

        if element == "a" {
            count_a += 1
        }

        if element == "e" {
            count_e += 1
        }

        if element == "i" {
            count_i += 1
        }

        if element == "o" {
            count_o += 1
        }

        if element == "u" {
            count_u += 1
        }
    }

    // Print statistic vowels
    fmt.Println("- Input string: ", str1)
    fmt.Println("- Count vowel [A]: ", count_a)
    fmt.Println("- Count vowel [E]: ", count_e)
    fmt.Println("- Count vowel [I]: ", count_i)
    fmt.Println("- Count vowel [O]: ", count_o)
    fmt.Println("- Count vowel [U]: ", count_u)
}
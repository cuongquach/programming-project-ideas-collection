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

	// List contains each reversed words
	var list_reversed_words []string

	for 0 <= len(list_each_words) {
		if len(list_each_words) == 0 {
			break
		}

		list_reversed_words = append(list_reversed_words, list_each_words[len(list_each_words)-1])

		// Remove last element in list_each_words
		list_each_words[len(list_each_words)-1] = ""
		list_each_words = list_each_words[:len(list_each_words)-1]
	}

	final_reversed_string := strings.Join(list_reversed_words, "")

	fmt.Println(final_reversed_string)
}
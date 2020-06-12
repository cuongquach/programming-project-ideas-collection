
package main
 
import (
    "fmt"
    "io/ioutil"
    "strings"
)

/*
- How many words ?
- How many characters ?
- How many sentences ?
- How many Paragraphs ?
- How many words count each ?
*/

func main() {
    // Read content in file cuongquach.txt
    data, err := ioutil.ReadFile("article.txt")
    if err != nil {
        fmt.Println(err)
    }
 
    // Convert data content to string
    content_file := string(data)
    content_file = strings.Replace(content_file, "\n", " ", -1)

    fmt.Println(content_file)

    // Count how many words in article
    amount_words := count_words(content_file)

    // Print summary
    fmt.Println("+ Words: ", amount_words)
}

func count_words(content string) int {
    

    // Remove useless characters like: , . -
    removed_character := strings.NewReplacer(",", " ",
                        ".", " ",
                        "-", " ")
    content = removed_character.Replace(content)
    fmt.Println(content)

    // Convert string to slice
    list_words := strings.Split(content, " ")

    for index, word := range list_words {

        if word == "" {
            fmt.Println(index)
            copy(list_words[index:], list_words[index+1:]) // Shift a[i+1:] left one index.
            fmt.Println(index)
            list_words[len(list_words)-1] = ""     // Erase last element (write zero value).
            list_words = list_words[:len(list_words)-1]     // Truncate slice.
            continue
        }
    }

    return len(list_words)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LetterCount struct {
    letter byte
    count int
}


func main() {
    file, err := os.Open("input.data")
    if err != nil {
        fmt.Println("Error reading file")
        fmt.Println(err)
    }
    defer file.Close()

    var niceStringsAmount = 0

    var fileReader *bufio.Reader = bufio.NewReader(file)

    counter := 0

    for {
        var word string = getNextString(fileReader)

        // If the word is empty, we reached the end of the file
        if word == "" {
            break;
        }
        counter++

        if isNiceString(word) {
            niceStringsAmount++;
        }

        fmt.Println("Counter:", counter)
        fmt.Println("Nice: ", niceStringsAmount)
    }

    fmt.Println("The amount of nice strings is: ", niceStringsAmount)
}

// Reads the next word from the file
func getNextString(fileReader *bufio.Reader) string {

    var word string
    var err error

    word, err = fileReader.ReadString('\n')

    if err != nil {
        return ""
    }

    // Remove the newline character
    word = word[:len(word) - 1]

    return word
}

// Verify if a word is a nice string
func isNiceString(word string) bool {
    var isWordNiceString bool
    var hasTwoLettersTwice bool
    var hasRepeatedLetter bool

    hasTwoLettersTwice = containsTwoLettersTwice(word)
    hasRepeatedLetter = containsRepeatedLetter(word)

    fmt.Println("......................")
    fmt.Println("Word:", word)
    fmt.Println("Has two letters twice:", hasTwoLettersTwice)
    fmt.Println("Has repeated letter:", hasRepeatedLetter)
    fmt.Println("......................")


    isWordNiceString = hasTwoLettersTwice && hasRepeatedLetter

    return isWordNiceString
}

func containsTwoLettersTwice(word string) bool {
    var group string
    var hasGroupTwice bool
    
    // Separate in groups of two
    for index := 0; index < len(word) - 1; index++ {
        group = word[index:index + 2]

        hasGroupTwice = strings.Count(word, group) >= 2 

        if hasGroupTwice{
            return true
        }
    }

    return false
}

// Verify if a word contains a letter which repeats with exactly
// one letter between them
func containsRepeatedLetter(word string) bool {
    var previousLetter string
    var nextLetter string
    var currentLetter string

    var isNiceString bool = false

    // Starts at the second letter and ends before the last
    for index := 1; index < len(word) - 1; index++ {
        previousLetter = string(word[index - 1])
        nextLetter = string(word[index + 1])
        currentLetter = string(word[index])

        // CONDITION:
        // It contains at least one letter which repeats with exactly
        // one letter between them, like "xyx", abcdefeghi (efe), or even aaa.

        if previousLetter == currentLetter {
            continue
        }

        isNiceString = previousLetter == nextLetter && previousLetter != currentLetter

        if isNiceString {
            break
        }
    }

    return isNiceString
}

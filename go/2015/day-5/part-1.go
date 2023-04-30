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

    for {
        var word string = getNextString(fileReader)

        // If the word is empty, we reached the end of the file
        if word == "" {
            break;
        }

        if isNiceString(word) {
            niceStringsAmount++;
        }
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
    var containsThreeVowels bool = countVowels(word) >= 3
    var containsDoubleLetter bool = false
    var containsForbiddenSequence bool = containsForbiddenSequence(word)

    var letterCounters []LetterCount = countLettersInRow(word)

    if len(letterCounters) > 0 {
        containsDoubleLetter = true
    }

    var isWordNiceString bool = containsThreeVowels && containsDoubleLetter && !containsForbiddenSequence

    return isWordNiceString
}

// Count the vowels in a word
func countVowels(word string) int {
    const VOWELS string = "aeiou"
    var vowelsCount int = 0

    // Loop through the word and count the vowels
    for index := 0; index < len(word); index++ {
        if strings.Contains(VOWELS, string(word[index])) {
            vowelsCount++
        }
    }

    return vowelsCount
}


func countLettersInRow(word string) []LetterCount {
    var counters []LetterCount

    for index := 0; index < len(word) - 1; index++ {
        var currentLetter byte = word[index]
        var nextLetter byte = word[index + 1]

        if currentLetter == nextLetter {
            if !containsLetter(counters, currentLetter) {
                counters = append(counters, LetterCount{currentLetter, 2})
            } else {
                // Find the counter with the current letter
                updateCounter(&counters, currentLetter)
            }
        }
    }

    return counters
}

func updateCounter(counters *[]LetterCount, letter byte) {
    for index := 0; index < len(*counters); index++ {
        if (*counters)[index].letter == letter {
            (*counters)[index].count++
            return
        }
    }
}

// Verify if a letter is already in the counters array
func containsLetter(counters []LetterCount, letter byte) bool {
    for index := 0; index < len(counters); index++ {
        if counters[index].letter == letter {
            return true
        }
    }

    return false
}

// Verify if the word contains any of the forbidden sequences
func containsForbiddenSequence(word string) bool {
    var hasFobbidenSequence bool = false
    var FORBIDDEN_SEQUENCES []string = []string{"ab", "cd", "pq", "xy"}


    for index := 1; index < len(word) && !hasFobbidenSequence; index++ {
        var currentLetter byte = word[index]
        var previousLetter byte = word[index - 1]

        for _, forbiddenSequence := range FORBIDDEN_SEQUENCES {
            var forbiddenSequenceFound = strings.Contains(forbiddenSequence, string(previousLetter) + string(currentLetter))

            if forbiddenSequenceFound {
                hasFobbidenSequence = true
                break
            }
        }
    }

    return hasFobbidenSequence
}

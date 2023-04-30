package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
    var secretKey string = readInput("input.data")
    var result string = ""
    var hash string = ""
    var input string = ""

    for index := 1; ; index++ {
        input = secretKey + fmt.Sprintf("%d", index)
        hash = generateMD5Key(input)

        if isAdventCoin(hash) {
            result = fmt.Sprint(index)
            break
        }
    }

    fmt.Println("=== AdventCoin found ===")
    fmt.Println("Number: ", result)
    fmt.Println("Input.: ", input)
    fmt.Println("Hash..: ", hash)
    fmt.Println("========================")
}

// Read the input file and return its contents as a string
func readInput(filename string) string {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return ""
    }
    defer file.Close()

    data, err := bufio.NewReader(file).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading file:", err)
        return ""
    }

    // Remove the newline character from the end of the string
    data = data[:len(data)-1]

    return data
}

// Generate MD5 key
func generateMD5Key(input string) string {
    // fmt.Println("Generating MD5 key for:", input)
    hash := md5.Sum([]byte(input))

    return hex.EncodeToString(hash[:])
}

// Verify if the key is an isAdventCoin,
// meaning that it starts with five zeroes
func isAdventCoin(key string) bool {
    return key[0:5] == "00000"
}

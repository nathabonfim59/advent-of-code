package main

// Disclaimer
// I don't know what I did wrong but this parallel
// version is significantly worse than the v1 one

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
)

type Data struct {
    index int
    input string
    hash string
}

func main() {
    var secretKey string = readInput("input.data")
    var input string = ""

    // Create a context
    ctx, cancel := context.WithCancel(context.Background())

    // Create a channel to communicate the result back to the main goroutine
    resultChan := make(chan Data)

    var numThreads int = runtime.NumCPU()

    for i := 0; i < numThreads; i++ {
        go func() {
            for index := 1; ; index++ {
                select {
                case <-ctx.Done():
                    return
                default:
                    input = secretKey + fmt.Sprintf("%d", index)

                    var hash string = generateMD5Key(input)

                    if isAdventCoin(hash) {
                        resultChan <- Data{index, input, hash}
                        cancel()
                        return
                    }
                }
            }
        }()
    }

    result := <-resultChan

    fmt.Println("=== AdventCoin found ===")
    fmt.Println("Number: ", result.index)
    fmt.Println("Input.: ", result.input)
    fmt.Println("Hash..: ", result.hash)
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
    hash := md5.Sum([]byte(input))

    return hex.EncodeToString(hash[:])
}

// Verify if the key is an AdventCoin,
// meaning that it starts with six zeroes
func isAdventCoin(key string) bool {
    return key[0:7] == "0000000"
}

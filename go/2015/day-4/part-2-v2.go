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
	"sync"
	"sync/atomic"
)

type Data struct {
    index int
    input string
    hash string
}

func main() {
    var secretKey string = readInput("input.data")
    var input string = ""
    var result Data = Data{0, "", ""}
    var resultFound int32 = 0

    var numThreads int = runtime.NumCPU()
    // var funcsPerThread int = 6

    // Create a context
	ctx, cancel := context.WithCancel(context.Background())

    // Create a wait group for waiting all threads to finish
    var wg sync.WaitGroup
    wg.Add(numThreads)

    // Add a mutex to protect access to the result variable
    var mu sync.Mutex

    for i := 0; i < numThreads; i++ {
        go func() {
            defer wg.Done()

            for index := 1; ; index++ {
                input = secretKey + fmt.Sprintf("%d", index)

                var input Data = Data{index, input, ""}

                select {
                    case <-ctx.Done():
                        return
                    default:
                        var hash string = generateMD5Key(input.input)

                        if isAdventCoin(hash) {
                            mu.Lock()
                            if atomic.CompareAndSwapInt32(&resultFound, 0, 1) {
                                result = Data{input.index, input.input, hash}
                                cancel()
                            }
                            mu.Unlock()
                            return
                        }
                }
            }
        }()
    }

    wg.Wait()

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
    // fmt.Println("Generating MD5 key for:", input)
    hash := md5.Sum([]byte(input))

    return hex.EncodeToString(hash[:])
}

// Verify if the key is an isAdventCoin,
// meaning that it starts with five zeroes
func isAdventCoin(key string) bool {
    return key[0:6] == "000000"
}

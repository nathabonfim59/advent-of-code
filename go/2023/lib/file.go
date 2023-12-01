package lib

import (
	"bufio"
	"os"
)

// Open a Reader for the given file
// to be read line by line.
func OpenFile(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil
}

// Read the next line from the given Reader.
func NextLine(reader *bufio.Reader) (string, error) {
    line, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }

    // Remove the last character ('\n')
    line = line[:len(line)-1]

    return line, nil
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w, "<h1>Hello World!</h1>")
}

// List all files under '/data' as JSON
func listData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*");

    dir, err := os.Open("./data")

    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("500 - Something bad happened!"))
        return;
    }

    files, err := dir.Readdir(-1)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"files":[`))

    for _, file := range files {
        // add the comma if it's not last item
        if file != files[len(files) - 1] {
            w.Write([]byte(fmt.Sprintf(`{"name": "%s"},`, file.Name())))
        } else {
            w.Write([]byte(fmt.Sprintf(`{"name": "%s"}`, file.Name())))
        }
    }

    w.Write([]byte("]}"))
}

func dataInfo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*");
    w.Header().Set("Content-Type", "application/json")

    filename := r.URL.Path

    filename = strings.TrimPrefix(filename, "/data/")
    filename = fmt.Sprintf("./data/%s", filename)

    fmt.Println("Requested file: ", filename)

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        fmt.Println(err)
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"error": true, "message": "404 - File not found!"}`))
        return
    }

    file, err := os.Open(filename)
    defer file.Close()
    
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"error": true, "message": "500 - Something bad happened!"`))
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"file":`))
    w.Write([]byte(fmt.Sprintf(`{"name": "%s"}`, file.Name())))
    
    w.Write([]byte(`,"steps": [`))

    scanner := bufio.NewScanner(file)
    isLastLine := false

    scanner.Scan()
    line := scanner.Text()

    for !isLastLine {
        // Verify if it's the last line
        if isLastLine {
            break
        }

        // The inputs comes as "{x, y}", trim brackets
        coords := strings.Split(line, ", ")
        coords[0] = strings.Trim(coords[0], "{")
        coords[3] = strings.Trim(coords[3], "}")

        w.Write([]byte(fmt.Sprintf(`{"coords":{"x": %s, "y": %s}, "entity": "%s", "direction": "%s"}`, coords[0], coords[1], coords[2], coords[3])))

        scanner.Scan()

        line = scanner.Text()
        isLastLine = line == "-"

        if !isLastLine {
            w.Write([]byte(","))
        }
    }

    w.Write([]byte("]}"))
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/data", listData)

    // dataInfo is a dynamic route, so it's necessary to use HandleFunc
    // directly instead of Handle
    http.HandleFunc("/data/", dataInfo)

    io.WriteString(os.Stdout, "Server started on port 3000\n");

    err := http.ListenAndServe(":3000", nil)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("Server closed under request")
    } else if err != nil {
        fmt.Println("Server closed unexpectedly")
        fmt.Println(err)
    }
}

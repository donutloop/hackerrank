package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the weightedUniformStrings function below.
func weightedUniformStrings(s string, queries []int32) []string {
    set := make(map[int32]int32)
    counter := make(map[int32]bool)

    if len(s) > 0 {
        weigth := int32(s[0]) - 96
        set[weigth] = set[weigth] + 1 
        counter[set[weigth]*weigth] = true
    }
    for i := 1; i < len(s); i++ {
        weigth := int32(s[i]) - 96
        if set[weigth] > 0 && s[i-1] != s[i] {
            set[weigth] = 0
        }
        set[weigth] = set[weigth] + 1 
        counter[set[weigth]*weigth] = true
    }
    result := make([]string, len(queries))
    var i int
    for _, q := range queries {
        _, ok := counter[q]
        if ok {
           result[i] = "Yes"     
        } else {
            result[i] = "No"
        }
        i++
    }
    fmt.Println(counter)
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    var queries []int32

    for i := 0; i < int(queriesCount); i++ {
        queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        queriesItem := int32(queriesItemTemp)
        queries = append(queries, queriesItem)
    }

    result := weightedUniformStrings(s, queries)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

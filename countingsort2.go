package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countingSort function below.
func countingSort(arr []int32) []int32 {
    count := make([]int32, 100)
    for i := 0; i < len(arr); i++ {
        count[arr[i]] = count[arr[i]] + 1
    }
    numbers := make([]int32, 0)
    for i := 0; i < len(count); i++ {
        n := count[i]
        if n == 0 {
            continue
        }
        for j := int32(0); j < n; j++ {
            numbers = append(numbers, int32(i))    
        }
    }    
    return numbers 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024 * 10)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := countingSort(arr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
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

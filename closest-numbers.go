package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the closestNumbers function below.
func closestNumbers(arr []int32) []int32 {
    closestDiff := int32(-1)
    pairs := make([]int32, 0)
    sort(arr)
    for j := 1; j < len(arr); j++ {
            diff := int32(math.Abs(float64(arr[j-1]-arr[j])))
            if closestDiff == -1 || diff < closestDiff {
                closestDiff = diff
                oldLength := len(pairs)
                pairs = make([]int32, 0 , oldLength)
            }
            if closestDiff == diff {
                pairs = append(pairs, arr[j-1])
                pairs = append(pairs, arr[j])
            }
    }
    
    return pairs
}

func sort(a []int32) {
    for j := 1; j < len(a); j++{
        key := a[j]
        i := j - 1
        for i >= 0 && a[i] > key {
            a[i + 1] = a[i]
            i = i - 1
        }
        a[i + 1] = key
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n) -1; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := closestNumbers(arr)

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

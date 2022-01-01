package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the missingNumbers function below.
func missingNumbers(arr []int32, brr []int32) []int32 {
    setB := make(map[int32]int32)
    for _, n := range brr {
        setB[n] = setB[n] + 1
    }
    setA := make(map[int32]int32)
    for _, n := range arr {
        setA[n] = setA[n] + 1
    }
    missing := make([]int32, 0)
    for n, counterB := range setB {
        counterA, ok := setA[n]
        if !ok {
            missing = append(missing, n)
        } else if counterB != counterA {
            missing = append(missing, n)
        }
    } 
    sort(missing)
    return missing
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

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    m := int32(mTemp)

    brrTemp := strings.Split(readLine(reader), " ")

    var brr []int32

    for i := 0; i < int(m); i++ {
        brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
        checkError(err)
        brrItem := int32(brrItemTemp)
        brr = append(brr, brrItem)
    }

    result := missingNumbers(arr, brr)

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

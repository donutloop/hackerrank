package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the insertionSort1 function below.
func insertionSort1(n int32, a []int32) {
    for j := 1; j < len(a); j++{
        key := a[j]
        i := j - 1
        for i >= 0 && a[i] > key {
            a[i + 1] = a[i]
            i = i - 1
            for k := 0; k < len(a); k++ {
                fmt.Print(a[k], " ")
            }
            fmt.Println()
        }
        a[i + 1] = key
    }
    for k := 0; k < len(a); k++ {
        fmt.Print(a[k], " ")
    }
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

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

    insertionSort1(n, arr)
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

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the beautifulTriplets function below.
func beautifulTriplets(d int, arr []int) int32 {
    count := int32(0)
    sort.Ints(arr)
    for k := len(arr) -1; k >= 0; k-- {
        for j := k - 1; j >= 0; j-- {
            dkj := arr[k] - arr[j]
            if d == dkj {
                for i := j - 1; i >= 0; i-- { 
                    dji := arr[j] - arr[i]
                    if d == dkj && dji == d {
                        count = count + 1
                    }
                }
            }
        }
    }
    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nd := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nd[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    dTemp, err := strconv.ParseInt(nd[1], 10, 64)
    checkError(err)
    d := int(dTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := beautifulTriplets(d, arr)

    fmt.Fprintf(writer, "%d\n", result)

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

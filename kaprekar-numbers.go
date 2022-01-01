package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {
    found := false
    for i := p; i <= q; i++ {
        if i == 1 {
            fmt.Print("1 ")
            found = true
            continue
        }
        v0 := int64(i) * int64(i)
        v1 := strconv.FormatInt(v0, 10)
        middle := len(v1) / 2
        if middle > 0 {
            left, err := strconv.Atoi(v1[:middle])
            if err != nil {
                panic("left side of string is not a number")
            }
            right, err := strconv.Atoi(v1[middle:])
            if err != nil {
                panic("right side of string is not number")
            }
            v2 := left + right
            if int32(v2) == i {
                found = true
                fmt.Print(v2, " ")
            }
        }
    }
    if !found {
        fmt.Println("INVALID RANGE")
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    p := int32(pTemp)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    kaprekarNumbers(p, q)
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

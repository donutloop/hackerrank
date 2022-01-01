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

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
    nkq := strings.Split(readLine(reader), " ")
    nTemp, _ := strconv.Atoi(nkq[0])
    b := float64(nTemp)
    kTemp, _ := strconv.Atoi(nkq[1])
    A := float64(kTemp)
    bH := 0.5 * b
    fmt.Printf("%.0f\n", math.Ceil(A/bH))
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}


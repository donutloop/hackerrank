package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the findPoint function below.
 */
func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
    rx := (2 * qx) - px
    ry := (2 * qy) - py
    return []int32{rx, ry}
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

    for nItr := 0; nItr < int(n); nItr++ {
        pxPyQxQy := strings.Split(readLine(reader), " ")

        pxTemp, err := strconv.ParseInt(pxPyQxQy[0], 10, 64)
        checkError(err)
        px := int32(pxTemp)

        pyTemp, err := strconv.ParseInt(pxPyQxQy[1], 10, 64)
        checkError(err)
        py := int32(pyTemp)

        qxTemp, err := strconv.ParseInt(pxPyQxQy[2], 10, 64)
        checkError(err)
        qx := int32(qxTemp)

        qyTemp, err := strconv.ParseInt(pxPyQxQy[3], 10, 64)
        checkError(err)
        qy := int32(qyTemp)

        result := findPoint(px, py, qx, qy)

        for resultItr, resultItem := range result {
            fmt.Fprintf(writer, "%d", resultItem)

            if resultItr != len(result) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        fmt.Fprintf(writer, "\n")
    }

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

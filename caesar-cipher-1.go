package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the caesarCipher function below.
func caesarCipher(ss string, k int32) string {
    newSentence := ""
    for _, s := range ss {
        if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) {
            if s >= 65 && s <= 90 {
                j := s + k
                for j > 90 {
                    j = (j - 90) + 64
                }
                newSentence += string(j)
            } else {
                j := s + k
                for j > 122 {
                    j = (j - 122) + 96
                }
                newSentence += string(j)
            }
        } else {
            newSentence += string(s)
        }
    }
    return newSentence
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
    if n == n {

    }
    s := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k)

    fmt.Fprintf(writer, "%s\n", result)

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

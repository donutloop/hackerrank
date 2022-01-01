package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "math"
)

// Complete the encryption function below.
func encryption(s string) string {
    rowLen := int(math.Floor(math.Sqrt(float64(len(s))))) +1
    colLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
    ss := make([]string, rowLen)
    for i := 0; i < rowLen; i++ {
       if colLen > len(s) -1 {
           colLen = len(s) 
       }
       raw := s[:colLen]
       s = s[colLen:]
       for j := 0; j < len(raw); j++ {
           ss[j] = ss[j] + string(raw[j])
       }  
    }

    var es string 
    for i := 0; i < len(ss); i++ {
        es = es + ss[i] + " "
    }
   
    return es
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

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

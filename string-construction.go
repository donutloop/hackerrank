package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the stringConstruction function below.
func stringConstruction(ss string) int32 {
    set := make(map[byte]struct{})
    p := "" 
    var cost int32
    for i := 0; i < len(ss); i++ {
      _, ok := set[ss[i]]
      if ok {
          last := ""
          for j := len(ss); j > i; j-- {
            if ok := strings.Contains(p, ss[i:j]); ok {
                last = ss[i:j]
            } else {
                break
            }
          }
          p += last
          continue
      }
      set[ss[i]] = struct{}{}
      p += string(ss[i])
      cost += 1
    }
    return cost
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := stringConstruction(s)

        fmt.Fprintf(writer, "%d\n", result)
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

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/big"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) *big.Int {
    if n == 1 {
        return big.NewInt(1)
    }
    if n == 0 {
        return big.NewInt(1)
    }
    sum := big.NewInt(1)
    for i := int64(1); i <= int64(n); i++ {
        sum = big.NewInt(0).Mul(sum, big.NewInt(i))  
    }
    return sum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    fmt.Println(extraLongFactorials(n))
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

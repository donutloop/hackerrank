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

// Complete the bonAppetit function below.
func bonAppetit(bill []int32, k int32, b int32) {
        sum := int32(0)
        for i := 0; i < len(bill); i++ {
            if int(k) == i {
                continue
            }

            sum = sum + bill[i]
        }
        diff := math.Abs(float64((sum / 2) - b))
        if diff == 0 {
            fmt.Println("Bon Appetit")
        } else {
            fmt.Println(diff)
        }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nk := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    billTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var bill []int32

    for i := 0; i < int(n); i++ {
        billItemTemp, err := strconv.ParseInt(billTemp[i], 10, 64)
        checkError(err)
        billItem := int32(billItemTemp)
        bill = append(bill, billItem)
    }

    bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    b := int32(bTemp)

    bonAppetit(bill, k, b)
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

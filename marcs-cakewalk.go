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

// Complete the marcsCakewalk function below.
func marcsCakewalk(calorie []int32) int64 {
    sort(calorie)
    var sum int64
    var j int
    for i := len(calorie)-1; i >= 0; i-- {
        sum = sum + (int64(calorie[i]) * int64(math.Pow(float64(2), float64(j))))
        j++
    }
    return sum
}

func sort(a []int32) {
    for j := 1; j < len(a); j++{
        key := a[j]
        i := j - 1
        for i >= 0 && a[i] > key {
            a[i + 1] = a[i]
            i = i - 1
        }
        a[i + 1] = key
    }
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

    calorieTemp := strings.Split(readLine(reader), " ")

    var calorie []int32

    for i := 0; i < int(n); i++ {
        calorieItemTemp, err := strconv.ParseInt(calorieTemp[i], 10, 64)
        checkError(err)
        calorieItem := int32(calorieItemTemp)
        calorie = append(calorie, calorieItem)
    }

    result := marcsCakewalk(calorie)

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

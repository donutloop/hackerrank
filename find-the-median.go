package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/rand"
)

// Complete the findMedian function below.
func findMedian(arr []int32) int32 {
    sortedArr := qsort(arr)
    return sortedArr[len(sortedArr)/2]
}

func qsort(a []int32) []int32 {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := rand.Int() % len(a)

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qsort(a[:left])
  qsort(a[left + 1:])


  return a
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024* 10)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := findMedian(arr)

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

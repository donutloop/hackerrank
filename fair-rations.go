package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the fairRations function below.
func fairRations(B []int32) string {
    i := 0
    count := 0
    max := int32(0)
    for _, b := range B {
       max = max + b
    }
    if max % 2 == 1 {
        return "NO"
    }
    for {
        evenCount := 0
        for _, b := range B {
            if b % 2 == 0 {
                evenCount = evenCount + 1
            }
        }
        if evenCount == len(B) {
            break
        }
        if B[i] % 2 == 0 {
            i = i + 1
            continue
        }
        B[i] = B[i] + 1
        if j := i - 1; j > 0 {
         if B[j] % 2 == 1 {
            i = i - 1
            B[i] =  B[i] + 1 
            count = count + 2
            continue
         }       
        }
        if j := i + 1; j < len(B) {
         if B[j] % 2 == 1 {
            i = i + 1
            B[i] = B[i] + 1    
            count = count + 2 
            continue
         }       
        }
         
        if j := i + 1; j < len(B) {
           i = i + 1
           B[i] = B[i] + 1    
           count = count + 2    
           continue    
        } else if  j := i - 1; j > 0 {
            i = i - 1
            B[i] =  B[i] + 1 
            count = count + 2
            continue
        }
        break 
    } 
    return strconv.Itoa(count)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)

    BTemp := strings.Split(readLine(reader), " ")

    var B []int32

    for i := 0; i < int(N); i++ {
        BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
        checkError(err)
        BItem := int32(BItemTemp)
        B = append(B, BItem)
    }

    result := fairRations(B)

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

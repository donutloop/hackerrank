package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumNumber function below.
func minimumNumber(n int32, password string) int32 {
    
    hasSpecialChar := func(c rune) bool {
        for _,sc := range []rune{'!','@','#','$','%','^','&','*','(', ')', '-', '+'} {
            if sc == c {
                return true
            }
        } 
        return false
    }
    
    lowerCase := int32(0)
    upperCase := int32(0)
    number := int32(0)
    specialChar := int32(0)
    for _, c := range password {
        if c >= 97 && c <= 122 {
            lowerCase = lowerCase + 1   
        } else if c >= 65 && c <= 90 {
            upperCase = upperCase + 1
        } else if c >= 48 && c <= 57 {
            number = number + 1    
        } else if hasSpecialChar(c) {
            specialChar = specialChar + 1   
        }
    }

    count := int32(0)
    if lowerCase == 0 {
        count = count + 1
    }
    if upperCase == 0 {
        count = count + 1
    }
    if number == 0 {
        count = count + 1
    }
    if specialChar == 0 {
        count = count + 1
    }

    if int32(len(password)) + count < 6 {
        count += 6 - (int32(len(password)) + count)
    }
    
    return count
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

    password := readLine(reader)

    answer := minimumNumber(n, password)

    fmt.Fprintf(writer, "%d\n", answer)

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

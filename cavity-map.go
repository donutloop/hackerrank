package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the cavityMap function below.
func cavityMap(grid []string) []string {
    if len(grid) == 1 {
        if len(grid[0]) == 1 {
            return grid
        }
    }
    newGrid := make([]string, 0)
    newGrid = append(newGrid, grid[0])
    for i := 1; i < len(grid) -1 ; i++ {
        s := string(grid[i][0])
        for j := 1; j < len(grid[i]) -1; j++ {
            right, _ := strconv.Atoi(string(grid[i][j-1]))
            up, _ := strconv.Atoi(string(grid[i -1][j]))
            current, _ := strconv.Atoi(string(grid[i][j]))
            down, _ := strconv.Atoi(string(grid[i + 1][j]))
            left, _ := strconv.Atoi(string(grid[i][j+1]))

            if current > right && current > left && current > down && current > up {
                s += "X"
            } else {
                s += string(grid[i][j])
            }
        }
        s += string(grid[i][len(grid[i])-1]) 
        newGrid = append(newGrid, s) 
    }
    newGrid = append(newGrid, grid[len(grid)-1])
    return newGrid
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

    var grid []string

    for i := 0; i < int(n); i++ {
        gridItem := readLine(reader)
        grid = append(grid, gridItem)
    }

    result := cavityMap(grid)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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

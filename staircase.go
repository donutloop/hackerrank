package main
import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    
    lenght, _ := strconv.Atoi(scanner.Text())
    
    for i:= 1; i <= lenght; i++{
        for k := lenght-i ; k > 0; k--{
            fmt.Print(" ")
        }
        for j:= i; j > 0; j--{
            fmt.Print("#")
        }
        fmt.Print("\n")
    }
}
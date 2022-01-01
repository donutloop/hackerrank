package main
import "fmt"
import "bufio"
import "os"
import "strings"


func main() {
 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanner.Scan()
    
    numbers := strings.Fields(scanner.Text())
    
    for i:= len(numbers) - 1; i >= 0; i--{
        fmt.Print(numbers[i])
        fmt.Print(" ")
    } 
}
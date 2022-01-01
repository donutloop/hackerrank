package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanner.Scan()
    numbersRaw := strings.Fields(scanner.Text())
    
    numbers := []int{}
    for _, v := range numbersRaw {
        n, _ := strconv.Atoi(v)
        numbers = append(numbers, n)
    }
    
    for {
        min := -1
        
        if len(numbers) == 0 {
            break
        }
        
        for _, newMin := range numbers {
            if newMin < min || min == -1 {
                min = newMin
            }
        }
        
        fmt.Println(len(numbers))
       
        for i := 0; i < len(numbers); i++{
            numbers[i] = numbers[i] - min 
            if numbers[i] == 0 || numbers[i] < 0  {
                numbers = append(numbers[:i], numbers[i+1:]...)
                i--
            }
        }    
    }
}
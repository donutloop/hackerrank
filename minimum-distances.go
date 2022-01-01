package main
import "fmt"
import "strconv"
import "strings"
import "bufio"
import "os"

func main() {
 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanner.Scan()
    raw := strings.Fields(scanner.Text())
    
    if scanner.Err() != nil {
        fmt.Print(scanner.Err())
        return 
    }
    
    vector := make([]int, 0,len(raw))
    for _,v := range raw {
        n, _ := strconv.Atoi(v)
        vector = append(vector, n)        
    }
  
    distances := make([]int, 0, len(vector) / 2)
    for i := 0; i < len(vector); i++{
        for j := i + 1; j < len(vector); j++{
            if vector[i] == vector[j] {
                  
                var d int 
                if i > j {
                    d = i - j
                }else if i < j {
                    d = j - i
                }else{
                    d = 0
                }
                
                distances = append(distances, d)
                break 
            }
        }   
    }
        
    min := -1                                        
    for _,v := range distances{
        if v < min || min == -1 {
            min = v
        } 
    }                                        
    fmt.Print(min)                                       
}
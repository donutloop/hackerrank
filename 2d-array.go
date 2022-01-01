package main
import "bufio"
import "os"
import "strings"
import "strconv"
import "fmt"


func main() {
 
    scanner := bufio.NewScanner(os.Stdin)    
    matrix := [][]int{}
    for scanner.Scan() {
        raw := strings.Fields(scanner.Text())
        
        vector := []int{}
        for _,v := range raw{
            number, _ := strconv.Atoi(v)
            vector = append(vector, number)
        }
        matrix = append(matrix, vector)
    }
   
    
    sums := []int{}
    for i := 1; i < len(matrix) - 1; i++ {    
        for j := 1; j < len(matrix[i]) - 1; j++ {    
            sum := matrix[i+1][j-1] + matrix[i+1][j] + matrix[i+1][j+1] + matrix[i][j] + matrix[i-1][j-1] + matrix[i-1][j] + matrix[i-1][j+1] 
         
            
            sums = append(sums, sum)
        }
    }
    
    
    max := sums[0] 
    for i:= 0; i < len(sums); i++ {
        if sums[i] > max{
            max = sums[i]
        } 
    }
    
        
    fmt.Print(max)    
}
package main
import "fmt"
import "bufio"
import "strconv"
import "strings"
import "os"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    tc, _ := strconv.Atoi(scanner.Text())
    
    for ; tc > 0; tc-- {
        scanner.Scan()
        raw := strings.Fields(scanner.Text())
        classProps := []int{}
        for _, v := range raw{
            prop,_ := strconv.Atoi(v)
            classProps = append(classProps, prop)  
        }
        
        scanner.Scan()
        raw = strings.Fields(scanner.Text())
        students := []int{} 
        for _, v := range raw {
            student,_ := strconv.Atoi(v)
            
            if student < 0 || student == 0 {
                students = append(students, student)    
            }         
        }
        
        if len(students) > classProps[1] || len(students) == classProps[1]{
            fmt.Println("NO")  
        }else{
            fmt.Println("YES")
        }
    }
    
}
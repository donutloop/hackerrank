package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    raw := strings.Fields(scanner.Text())
    props := []int{}
    for _,v := range raw {
        n,_ := strconv.Atoi(v)
        props = append(props, n)        
    }
    
    if err := scanner.Err(); err != nil {
            fmt.Println(err)
    }
    
    x1 := props[0]
    v1 := props[1]
    x2 := props[2]
    v2 := props[3]
   
    d := distance(x1, x2) 
    dx := 0
    for dx < d {
        x2 = x2 + v2
        x1 = x1 + v1
        dx = distance(x1, x2)
        if x1 == x2 {
            fmt.Print("YES")
            return
        }
    }
    
    fmt.Print("NO")
}

func distance(x1, x2 int) (d int){
    if x1 > x2 {
        d = x1 - x2
    }else if x2 > x1 {
        d = x2 - x1
    }
    return 
}

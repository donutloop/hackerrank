package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    raw := strings.Fields(scanner.Text())
    if scanner.Err() != nil {
        fmt.Print(scanner.Err())
        return
    }
    props := []int{}
    for _,v := range raw {
        n,_ := strconv.Atoi(v)
        props = append(props, n)
    }
    
    x := props[0]
    end := props[1]
    m := props[2]
    
    var count int
    for x <= end {
        xr := reverse(x)
        if (x - xr) % m == 0 {
            count++
        } 
        x++
    }
    
    fmt.Print(count)
}

func reverse(x int) int {
    rev :=  0;
    for x != 0 {
        rev = rev*10 + x%10;
        x = x/10;
    }
    return rev;
}
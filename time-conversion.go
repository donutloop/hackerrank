package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    
    dateParts := strings.Split(scanner.Text(), ":")
    
    var mode int
    
    if strings.Contains(dateParts[2], "PM"){
        if dateParts[0] != "12" {
            mode = 12    
        }
    }else {
        if dateParts[0] == "12"{
            mode = -12 
        }
    }
    
    hours, _ := strconv.Atoi(dateParts[0])
    h := strconv.Itoa(hours + mode)
    
    if len(h) == 1{
        h = "0"  + h
    }
    
    fmt.Print(h + ":" + dateParts[1] + ":" + dateParts[2][:len(dateParts[2])- 2 ])  
}
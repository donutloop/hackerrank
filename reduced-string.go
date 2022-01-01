package main
import "fmt"
import "bufio"
import "os"

func main() {
 
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    
    text := scanner.Text()
    var found bool
    for {
        if text, found = clean(text); found {
            clean(text)
        }else{
            break 
        }
    }
    
    if text == ""{
        text = "Empty String"
    }
    
    fmt.Print(text)
}

func clean(dirtyText string) (text string ,found bool){
    for i:= 0; i < len(dirtyText); i++{
            if ((i + 1) < len(dirtyText) && dirtyText[i + 1] == dirtyText[i] ) {
                i++
                found = true 
                continue
            }
            text += string(dirtyText[i])
        }
    return 
}

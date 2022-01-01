package main

import "bufio"
import "os"
import "strconv"
import "fmt"
import "math"

func main(){
    
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    mealCost, _ := strconv.ParseFloat(scanner.Text(), 64)
    
    scanner.Scan()
    tipP, _ := strconv.ParseFloat(scanner.Text(), 64)
    
    scanner.Scan()
    taxP, _ := strconv.ParseFloat(scanner.Text(), 64)
    
    tip := mealCost * (tipP / 100.0)
    tax := mealCost * (taxP/ 100.0)
    
    finalCost := mealCost + tip + tax 
    
    finalCost = math.Floor(finalCost + .5)
    
    fmt.Printf("The total meal cost is %v dollars.", finalCost)
}
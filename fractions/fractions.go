package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
)

func main() {
    percentages := os.Args[1:]
    pLen := len(percentages)
    if pLen != 2 {
        fmt.Printf("Invalid number of percentages:  %d\n", pLen)
        return
    }
    lowerP, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
        fmt.Printf("Error:  %s\n", err.Error())
        return
    }
    upperP, err := strconv.ParseFloat(os.Args[2], 64)
    if err != nil {
        fmt.Printf("Error:  %s\n", err.Error())
        return
    }
    if upperP <= lowerP {
        fmt.Println("Error:  first percentage cannot be greater than or equal to second")
        return
    }
    if lowerP < 0 || 100 < upperP {
        fmt.Println("Values must be between 0 and 100")
        return
    }
    maxDivisor := int(math.Ceil(100.0 / (upperP - lowerP)))
    var divisor int
    var dividend int
    for divisor = 1; divisor <= maxDivisor; divisor++ {
        fractionFound := false
        for dividend = 0; dividend <= divisor; dividend++ {
            percent := float64(dividend)/float64(divisor)*100.0
            if lowerP < percent && percent <= upperP {
                fractionFound = true
                break
            } else if upperP < percent {
                continue
            }
        }
        if fractionFound {
            break
        }
    }
    fmt.Printf("Fraction:  %d/%d\n", dividend, divisor)
}

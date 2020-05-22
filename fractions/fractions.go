package main

import (
    "flag"
    "fmt"
    "math"
    "strconv"
)

func printHelp() {
    fmt.Println("Usage:  fractions <lower percent> <upper percent>")
}

func main() {

    helpShortPtr := flag.Bool("h", false, "help menu")
    helpLongPtr := flag.Bool("help", false, "help menu")
    flag.Parse()
    if *helpShortPtr || *helpLongPtr {
        printHelp()
        return
    }

    args := flag.Args()
    pLen := len(args)
    if pLen != 2 {
        fmt.Printf("Invalid number of percentages:  %d\n", pLen)
        return
    }
    lowerP, err := strconv.ParseFloat(args[0], 64)
    if err != nil {
        fmt.Printf("Error:  %s\n", err.Error())
        return
    }
    upperP, err := strconv.ParseFloat(args[1], 64)
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

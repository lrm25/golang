package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {

    // get a real random-ish number
    rand.Seed(time.Now().UnixNano())

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Pick a number between 1 and 10:  ")
    input, _ := reader.ReadString('\n')
    trimmedInput := strings.TrimSuffix(input, "\n")

    inputNumber, err := strconv.Atoi(trimmedInput)
    if err != nil {
        fmt.Printf("Either '%s' isn't a number ", trimmedInput)
        fmt.Printf("or you're trying to be too clever.  ")
        fmt.Println("No decimals, pi, e, any of that shit.")
        return
    }
    if inputNumber < 1 || 10 < inputNumber {
        fmt.Println("I said between 1 and 10!")
        return
    }
    randomNumber := rand.Intn(10) + 1

    fmt.Printf("The number is %d\n", randomNumber)
    fmt.Printf("You guessed %d\n", inputNumber)
    if inputNumber != randomNumber {
        fmt.Println("You LOSE!  HA!")
        return
    }
    fmt.Println("Fine.  You win.")
}

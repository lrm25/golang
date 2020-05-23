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

    inputNumber, _ := strconv.Atoi(trimmedInput)
    randomNumber := rand.Intn(10) + 1

    fmt.Printf("The number is %d\n", randomNumber)
    fmt.Printf("You guessed %d\n", inputNumber)
    if inputNumber != randomNumber {
        fmt.Println("You LOSE!  HA!")
        return
    }
    fmt.Println("Fine.  You win.")
}

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func roll() int {
    dieOne := rand.Intn(5) + 1
    dieTwo := rand.Intn(5) + 1
    roll := dieOne + dieTwo
    fmt.Printf("Roll:  %d\n", roll)
    return roll
}

func playRound(bet int) int {
    rollOne := roll()
    if rollOne == 7 || rollOne == 11 {
        fmt.Println("You win!")
        return bet * 2
    }
    if rollOne == 2 || rollOne == 3 || rollOne == 12 {
        fmt.Println("You lose!")
        return 0
    }
    for {
        roll := roll()
        if roll == rollOne {
            fmt.Println("You win!");
            return bet * 2
        }
        if roll == 7 {
            fmt.Println("You lose!");
            return 0
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    money := 10
    fmt.Printf("Betting %d dollars\n", money)
    moneyLeft := playRound(money)
    fmt.Printf("You now have %d dollars\n", moneyLeft)
}

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

func main() {
    rand.Seed(time.Now().UnixNano())
    rollOne := roll()
    if rollOne == 7 || rollOne == 11 {
        fmt.Println("You win!")
        return
    }
    if rollOne == 2 || rollOne == 3 || rollOne == 12 {
        fmt.Println("You lose!")
        return
    }
    for {
        roll := roll()
        if roll == rollOne {
            fmt.Println("You win!");
            return
        }
        if roll == 7 {
            fmt.Println("You lose!");
            return
        }
    }
}

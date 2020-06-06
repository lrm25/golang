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

    reader := bufio.NewReader(os.Stdin)
    money := 100
    fmt.Printf("You have %d dollars\n", money)

    rand.Seed(time.Now().UnixNano())
    for 0 < money {
        fmt.Printf("Play round (y/n): ")
        input, _ := reader.ReadString('\n')
        answer := strings.TrimSpace(input)
        if answer == "y" || answer == "Y" {
            for {
                fmt.Printf("Place bet: ")
                rawBetInput, _ := reader.ReadString('\n')
                betString := strings.TrimSpace(rawBetInput)
                bet, err := strconv.Atoi(betString)
                if err != nil {
                    fmt.Println("Integers please")
                    continue
                } else if bet <= 0 || money < bet {
                    fmt.Println("More than 0, less than or what you have")
                    continue
                }
                fmt.Printf("Betting %d dollars\n", bet)
                money -= bet
                money += playRound(bet)
                fmt.Printf("You now have %d dollars\n", money)
                break
            }
        } else if answer == "n" || answer == "N" {
            fmt.Println("Wuss")
            break
        } else {
            fmt.Println("huh?")
        }
    }
    if money == 0 {
        fmt.Println("Time to go home.")
    }
}

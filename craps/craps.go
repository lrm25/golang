package main

import (
    "bufio"
    "flag"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

var debugMode bool = false

func roll() int {
    var roll int
    if debugMode {
        fmt.Printf("Enter roll: ")
        reader := bufio.NewReader(os.Stdin)
        rollInput, _ := reader.ReadString('\n')
        rollString := strings.TrimSpace(rollInput)
        roll, _ = strconv.Atoi(rollString)
    } else {
        dieOne := rand.Intn(5) + 1
        dieTwo := rand.Intn(5) + 1
        roll := dieOne + dieTwo
        fmt.Printf("Roll:  %d\n", roll)
    }
    return roll
}

func playRound(bet int, passBet bool) int {

    var pass bool
    draw := false
    complete := false

    rollOne := roll()
    if rollOne == 7 || rollOne == 11 {
        pass = true
        complete = true
    }
    if rollOne == 2 || rollOne == 3 {
        pass = false
        complete = true
    }
    if rollOne == 12 {
        pass = false
        draw = true
        complete = true
    }
    for !complete {
        roll := roll()
        if roll == rollOne {
            pass = true
            complete = true
        }
        if roll == 7 {
            pass = false
            complete = true
        }
    }
    if pass == passBet {
        if draw {
            fmt.Println("Draw")
            return bet
        } else {
            fmt.Println("You win!")
            return bet * 2
        }
    } else {
        fmt.Println("You lose!")
        return 0
    }
}

func main() {

    debugPtr := flag.Bool("debug", false, "debug mode")
    flag.Parse()
    if *debugPtr {
        debugMode = true
    }

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
                var pass bool
                for {
                    fmt.Printf("(p)ass/(d)on't: ")
                    passInput, _ := reader.ReadString('\n')
                    passString := strings.TrimSpace(passInput)
                    if passString == "p" || passString == "d" {
                        if passString == "p" {
                            pass = true
                        } else {
                            pass = false
                        }
                        break
                    }
                    fmt.Println("p or d")
                }
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
                money += playRound(bet, pass)
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

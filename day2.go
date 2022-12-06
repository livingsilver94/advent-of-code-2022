package main

import (
    "bufio"
    "errors"
    "io"
    "log"
    "os"
    "strings"
)

type MyMove rune

const (
    MyRock    MyMove = 'X'
    MyPaper   MyMove = 'Y'
    MyScissor MyMove = 'Z'
)

func (m MyMove) Points() int {
    switch m {
    case MyRock:
        return 1
    case MyPaper:
        return 2
    case MyScissor:
        return 3
    }
    panic(m)
}


type TheirMove rune

const (
    TheirRock    TheirMove = 'A'
    TheirPaper   TheirMove = 'B'
    TheirScissor TheirMove = 'C'
)

type Outcome rune

const (
    Victory Outcome = 'Z'
    Draw    Outcome = 'Y'
    Defeat  Outcome = 'X'
)

func (o Outcome) Points() int {
    switch o {
    case Victory:
        return 6
    case Draw:
        return 3
    case Defeat:
        return 0
    }
     panic(o)
}

func MyMoveTo(o Outcome, their TheirMove) MyMove {
    switch o {
    case Draw:
        return  MyMove(their + 23) // ASCII distance.
    case Victory:
        switch their {
        case TheirRock:
            return MyPaper
        case TheirPaper:
            return MyScissor
        case TheirScissor:
            return MyRock
        }
    case Defeat:
        switch their {
        case TheirRock:
            return MyScissor
        case TheirPaper:
            return MyRock
        case TheirScissor:
            return MyPaper
        }
    }
    panic(o)
}


func GetOutcome(my MyMove, their TheirMove) Outcome {
    var th MyMove = MyMove(their + 23) // ASCII distance.
    if my == th {
        return Draw
    }
    switch my {
    case MyRock:
        if th == MyScissor {
            return Victory
        }
        return Defeat
    case MyPaper:
        if th == MyRock {
            return Victory
        }
        return Defeat
    case MyScissor:
        if th == MyPaper {
            return Victory
        }
        return Defeat
    }
    panic(my)
}

func Score(my MyMove, out Outcome) int {
    return my.Points() + out.Points()
}

func main() {
    var (
        reader = bufio.NewReader(os.Stdin)
        part1score = 0
        part2score = 0
    )

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if errors.Is(err, io.EOF) {
                log.Println(part1score)
                log.Println(part2score)
                return
            }
            log.Fatal(err)
        }
        line = line[:len(line)-1]
        if line == "" {
            continue
        } else {
            moves := strings.Split(line, " ")
            myMove := MyMove(moves[1][0])
            thMove := TheirMove(moves[0][0])

            part1outcome := GetOutcome(myMove, thMove)
            part1score += Score(MyMove(moves[1][0]), part1outcome)

            expOutcome := Outcome(moves[1][0])
            expMyMove := MyMoveTo(expOutcome, thMove)
            part2score += Score(expMyMove, expOutcome)
        }
    }
}

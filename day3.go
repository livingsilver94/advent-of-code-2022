package main

import (
    "bufio"
    "errors"
    "io"
    "log"
    "os"
)

type Item rune

func (i Item) Priority() int {
    switch {
    case i >= 'a' && i <= 'z':
        return int(i) - 96 // ASCII distance.
    default:
        return int(i) - 38 // ASCII distance.
    }
}

type Compartment map[Item]struct{}

func (c Compartment) Insert(item Item) {
    c[item] = struct{}{}
}

func (c Compartment) Contains(item Item) bool {
    _, ok := c[item]
    return ok
}

func (c Compartment) Intersect(c2 Compartment) Compartment {
    out := make(Compartment)
    for item, _ := range c {
        if !c2.Contains(item) {
            continue
        }
        out.Insert(item)
    }
    return out
}

func (c Compartment) Union(c2 Compartment) Compartment {
    out := make(Compartment)
    for item, _ := range c {
        out[item] = struct{}{}
    }
    for item, _ := range c2 {
        out[item] = struct{}{}
    }
    return out
}

func ParseCompartments(s string) (Compartment, Compartment) {
    c1 := make(Compartment)
    c2 := make(Compartment)
    for _, item := range s[:len(s)/2] {
        c1.Insert(Item(item))
    }
    for _, item := range s[len(s)/2:] {
        c2.Insert(Item(item))
    }
    return c1, c2
}

func main() {
    var (
        reader = bufio.NewReader(os.Stdin)
        priorityPart1 = 0

        lineCount = 1
        rucksack1 Compartment
        rucksack2 Compartment
        rucksack3 Compartment
        priorityPart2 = 0
    )

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if errors.Is(err, io.EOF) {
                log.Println(priorityPart1)
                log.Println(priorityPart2)
                return
            }
            log.Fatal(err)
        }
        line = line[:len(line)-1]
        if line == "" {
            continue
        } else {
            // Part 1.
            comp1, comp2 := ParseCompartments(line)
            inter := comp1.Intersect(comp2)
            for item, _ := range inter {
                priorityPart1 += item.Priority()
            }

            // Part 2.
            switch lineCount {
            case 1:
                rucksack1 = comp1.Union(comp2)
                lineCount += 1
            case 2:
                rucksack2 = comp1.Union(comp2)
                lineCount += 1
            case 3:
                rucksack3 = comp1.Union(comp2)
                badges := rucksack1.Intersect(rucksack2).Intersect(rucksack3)
                for item, _ := range badges {
                    priorityPart2 += item.Priority()
                }
                lineCount = 1
            }
        }
    }
}

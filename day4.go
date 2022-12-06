package main

import (
    "bufio"
    "errors"
    "io"
    "log"
    "os"
    "strconv"
    "strings"
)

type Range struct {
    From int
    To   int
}

func NewRange(from, to int) Range {
    return Range {
        From: from,
        To: to,
    }
}

func (r Range) Length() int {
    return r.To - r.From
}

func (r Range) Contains(r2 Range) bool {
    return r.To >= r2.To && r.From <= r2.From
}

func (r Range) Overlaps(r2 Range) bool {
    return r2.From >= r.From && r2.From <= r.To
}

func ParsePairs(s string) (Range, Range) {
    ranges := make([]Range, 2)
    for i, pairString := range strings.Split(s, ",") {
        rangeString := strings.Split(pairString, "-")
        from, err := strconv.Atoi(rangeString[0])
        if err != nil {
            log.Fatal(err)
        }
        to, err := strconv.Atoi(rangeString[1])
        if err != nil {
            log.Fatal(err)
        }
        ranges[i] = NewRange(from, to)
    }
    return ranges[0], ranges[1]
}

func main() {
    var (
        reader = bufio.NewReader(os.Stdin)

        containCount = 0
        overlapCount = 0
    )

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if errors.Is(err, io.EOF) {
               log.Println(containCount)
               log.Println(overlapCount)
                return
            }
            log.Fatal(err)
        }
        line = line[:len(line)-1]
        if line == "" {
            continue
        } else {
            // Part 1.
            range1, range2 := ParsePairs(line)
            if range1.Contains(range2) || range2.Contains(range1) {
                containCount += 1
            }

            if range1.Overlaps(range2) || range2.Overlaps(range1) {
                overlapCount += 1
            }
        }
    }
}

package main

import (
    "bufio"
    "errors"
    "io"
    "log"
    "os"
    //"strconv"
    //"strings"
)

const (
    Part1WindowSize = 4
    Part2WindowSize = 14
)

func UniqueString(s string) bool {
    for i, ch1 := range s {
        if i+1 == len(s) {
            return true
        }
        for _, ch2 := range s[i+1:] {
            if ch1 == ch2 {
                return false
            }
        }
    }
    return true
}

func IndexAfterMarker(s string, markerLen int) int {
    for i := 0; i + markerLen < len(s); i+=1 {
        marker := s[i:i+markerLen]
        if !UniqueString(marker) {
            continue
        }
        return i + markerLen
    }
    return -1
}

func main() {
    var (
        reader = bufio.NewReader(os.Stdin)
        
        part1AfterMarker int
        part2AfterMarker int
    )

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if errors.Is(err, io.EOF) {
               log.Println(part1AfterMarker)
               log.Println(part2AfterMarker)
                return
            }
            log.Fatal(err)
        }
        line = line[:len(line)-1]
        if line == "" {
            continue
        } else {
            part1AfterMarker = IndexAfterMarker(line, Part1WindowSize)
            part2AfterMarker = IndexAfterMarker(line, Part2WindowSize)
        }
    }
}

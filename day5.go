package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

// We may create a fixed size stack type by reading the
// maximium size, as it cane be calculated by findi "\n\n"
// in the input file, but since we have no performance
// constraint let's just keep it simple.

type Stack []rune

func NewStack() Stack {
	return make(Stack, 0, 10)
}

func (s *Stack) Clone() Stack {
    ret := make(Stack, len(*s))
    copy(ret, *s)
    return ret
}

func (s *Stack) Push(i rune) {
	*s = append(*s, i)
}

func (s *Stack) PushN(i []rune) {
    *s = append(*s, i...)
}

func (s *Stack) Pop() (rune, bool) {
	length := len(*s)
	if length == 0 {
		return 0, false
	}
	i := (*s)[length-1]
	*s = (*s)[:length-1]
	return i, true
}

func (s *Stack) PopN(n int) ([]rune, bool) {
    length := len(*s)
	if length < n {
        panic(n)
		return nil, false
	}
	ret := (*s)[length-n:]
    *s = (*s)[:length-n]
    return ret, true
}

func (s *Stack) Reverse() {
	s2 := make(Stack, 0, len(*s))
	for _ = range *s {
		if val, ok := s.Pop(); ok {
			s2.Push(val)
		}
	}
	*s = s2
}

type Move struct {
	FromStack int
	ToStack   int
	Times     int
}

func NewMove(from, to, times int) Move {
	return Move{
		FromStack: from,
		ToStack:   to,
		Times:     times,
	}
}

func TextBlocks(r io.Reader) (string, string) {
	text, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	block1, block2, ok := strings.Cut(string(text), "\n\n")
	if !ok {
		panic(ok)
	}
	return block1, block2
}

func ParseStacks(s string) []Stack {
	reader := bufio.NewReader(strings.NewReader(s))
	stacks := make([]Stack, 0, 10)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}
			break
		}
		line = line[:len(line)-1]
		stackIndex := -1
		for i := 1; i < len(line); i += 4 {
			stackIndex += 1
			if len(stacks) <= stackIndex {
				stacks = append(stacks, NewStack())
			}
			crate := rune(line[i])
			if unicode.IsSpace(crate) {
				continue
			}
			stacks[stackIndex].Push(crate)
		}
	}
	for i := range stacks {
		stacks[i].Reverse()
	}
	return stacks
}

func ParseMoves(s string) []Move {
	reader := bufio.NewReader(strings.NewReader(s))
	moves := make([]Move, 0, 100)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}
			break
		}
		line = line[:len(line)-1]
		var (
			times     int
			fromStack int
			toStack   int
		)
		fmt.Sscanf(line, "move %d from %d to %d", &times, &fromStack, &toStack)
		moves = append(moves, NewMove(fromStack, toStack, times))
	}
	return moves
}

func StacksString(stacks []Stack) string {
    var b strings.Builder
	for _, stack := range stacks {
		val, ok := stack.Pop()
		if !ok {
			continue
		}
		b.WriteRune(val)
	}
	return b.String()
}

func main() {
	stacksBlock, movesBlock := TextBlocks(bufio.NewReader(os.Stdin))
	stacks := ParseStacks(stacksBlock)
	moves := ParseMoves(movesBlock)

    // Part 1.
    part1Stacks := make([]Stack, len(stacks))
    for i := range stacks {
        part1Stacks[i] = stacks[i].Clone()
    }
	for _, move := range moves {
		from := &part1Stacks[move.FromStack-1]
		to := &part1Stacks[move.ToStack-1]
		for i := 0; i < move.Times; i++ {
			if val, ok := from.Pop(); ok {
				to.Push(val)
			}
		}
	}
	log.Println(StacksString(part1Stacks))

    // Part 2.
    part2Stacks := make([]Stack, len(stacks))
    for i := range stacks {
        part2Stacks[i] = stacks[i].Clone()
    }
    for _, move := range moves {
		from := &part2Stacks[move.FromStack-1]
		to := &part2Stacks[move.ToStack-1]
        if val, ok := from.PopN(move.Times); ok {
            to.PushN(val)
		}
	}
	log.Println(StacksString(part2Stacks))
}

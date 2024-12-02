package pkg

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func Solve1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, seq := range strings.Split(line, ",") {
			sum += hash(seq)
		}
	}
	return sum, nil
}

func Solve2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	b := boxes{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, seq := range strings.Split(line, ",") {
			b.process(seq)
		}
	}
	return b.power(), nil
}

func hash(input string) int {
	current := 0
	for _, c := range input {
		current = ((current + int(c)) * 17) % 256
	}
	return current
}

type lens struct {
	label string
	focal int
}

type boxes map[int][]lens

func (b boxes) process(seq string) {
	re := regexp.MustCompile(`(.+)(=|-)(.*)`)
	matches := re.FindStringSubmatch(seq)
	if len(matches) != 3 && len(matches) != 4 {
		panic(fmt.Sprint("invalid sequence", seq))
	}
	label := matches[1]
	op := matches[2]
	box := hash(label)
	if op == "-" {
		b.remove(box, label)
	} else {
		b.add(box, lens{label, int(matches[3][0]) - '0'})
	}
}

func (b boxes) remove(box int, label string) {
	lenses := b[box]
	for i, l := range lenses {
		if l.label == label {
			b[box] = append(lenses[:i], lenses[i+1:]...)
			if len(b[box]) == 0 {
				delete(b, box)
			}
			return
		}
	}
}

func (b boxes) add(box int, newLens lens) {
	lenses := b[box]
	for i, l := range lenses {
		if l.label == newLens.label {
			b[box][i] = newLens
			return
		}
	}
	b[box] = append(lenses, newLens)
}

func (b boxes) power() int {
	total := 0
	for boxNum := range b {
		for i, l := range b[boxNum] {
			total += (boxNum + 1) * (i + 1) * l.focal
		}
	}
	return total
}

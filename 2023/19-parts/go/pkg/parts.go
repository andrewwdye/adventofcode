package pkg

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Rule struct {
	Category    byte
	Comparator  byte
	Value       int
	Destination string
}

func (r Rule) String() string {
	return fmt.Sprintf("%c%c%d:%s", r.Category, r.Comparator, r.Value, r.Destination)
}

type Workflow struct {
	Rules []Rule
}

type Part map[byte]int

var workflowRe = regexp.MustCompile(`([a-z]+)\{(.+)\}`)
var ruleRe = regexp.MustCompile(`([xmas])([<>])([0-9]+):([a-zA-Z]+)`)

func getWorkflows(workflowLines []string) map[string]Workflow {
	workflows := make(map[string]Workflow, len(workflowLines))
	for _, line := range workflowLines {
		matches := workflowRe.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic(fmt.Sprint("invalid workflow line: ", line))
		}
		name := matches[1]
		ruleStrs := strings.Split(matches[2], ",")
		w := Workflow{
			Rules: make([]Rule, 0, len(ruleStrs)),
		}
		for _, r := range ruleStrs[:len(ruleStrs)-1] {
			matches := ruleRe.FindStringSubmatch(r)
			if len(matches) != 5 {
				panic(fmt.Sprint("invalid rule: ", r))
			}
			rule := Rule{
				Category:    matches[1][0],
				Comparator:  matches[2][0],
				Destination: matches[4],
			}
			value, err := strconv.Atoi(matches[3])
			if err != nil {
				panic(err)
			}
			rule.Value = value
			w.Rules = append(w.Rules, rule)
		}
		w.Rules = append(w.Rules, Rule{
			Destination: ruleStrs[len(ruleStrs)-1],
		})
		workflows[name] = w
	}
	return workflows
}

func getParts(ratingLines []string) []Part {
	re := regexp.MustCompile(`\{x=([0-9]+),m=([0-9]+),a=([0-9]+),s=([0-9]+)\}`)
	parts := make([]Part, 0, len(ratingLines))
	for _, line := range ratingLines {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 5 {
			panic(fmt.Sprint("invalid rating line", line))
		}
		parts = append(parts, Part{
			'x': lo.Must(strconv.Atoi(matches[1])),
			'm': lo.Must(strconv.Atoi(matches[2])),
			'a': lo.Must(strconv.Atoi(matches[3])),
			's': lo.Must(strconv.Atoi(matches[4])),
		})
	}
	return parts
}

func process(workflowLines, ratingLines []string) int {
	workflows := getWorkflows(workflowLines)
	parts := getParts(ratingLines)

	sum := 0
	for _, part := range parts {
		if processPart(workflows, part) {
			for _, v := range part {
				sum += v
			}
		}
	}
	return sum
}

func processPart(workflows map[string]Workflow, part Part) bool {
	w := workflows["in"]
	for {
		var dest string
		for _, rule := range w.Rules {
			c := part[rule.Category]
			v := rule.Value
			if (rule.Comparator == '<' && c < v) || (rule.Comparator == '>' && c > v) || rule.Comparator == byte(0) {
				dest = rule.Destination
			} else {
				continue
			}
			if dest == "A" {
				return true
			}
			if dest == "R" {
				return false
			}
			w = workflows[dest]
			break
		}
	}
}

type Ranges struct {
	min map[byte]int
	max map[byte]int
}

func (r Ranges) Copy() Ranges {
	other := Ranges{
		make(map[byte]int, len(r.min)),
		make(map[byte]int, len(r.max)),
	}
	for k, v := range r.min {
		other.min[k] = v
	}
	for k, v := range r.max {
		other.max[k] = v
	}
	return other
}

func (r Ranges) Ways() int {
	ways := 1
	for k := range r.min {
		span := r.max[k] - r.min[k] + 1
		if span > 0 {
			ways *= span
		}
	}
	return ways
}

func NewRanges() Ranges {
	return Ranges{
		min: map[byte]int{
			'x': 1,
			'm': 1,
			'a': 1,
			's': 1,
		},
		max: map[byte]int{
			'x': 4000,
			'm': 4000,
			'a': 4000,
			's': 4000,
		},
	}
}

type Elem struct {
	workflow Workflow
	ranges   Ranges
}

func ways(workflowLines []string) int {
	workflows := getWorkflows(workflowLines)

	elems := []Elem{
		{
			workflows["in"],
			NewRanges(),
		},
	}
	ways := 0
	for len(elems) > 0 {
		elem := elems[0]
		elems = elems[1:]
		remaining := elem.ranges.Copy()
		for _, rule := range elem.workflow.Rules {
			thisRange := remaining.Copy()
			if rule.Comparator == '<' {
				if rule.Value-1 < thisRange.max[rule.Category] {
					thisRange.max[rule.Category] = rule.Value - 1
				}
				if rule.Value > remaining.min[rule.Category] {
					remaining.min[rule.Category] = rule.Value
				}
			} else if rule.Comparator == '>' {
				if rule.Value+1 > thisRange.min[rule.Category] {
					thisRange.min[rule.Category] = rule.Value + 1
				}
				if rule.Value < remaining.max[rule.Category] {
					remaining.max[rule.Category] = rule.Value
				}
			}
			switch rule.Destination {
			case "R":
				continue
			case "A":
				ways += thisRange.Ways()
			}
			elems = append(elems, Elem{
				workflows[rule.Destination],
				thisRange,
			})
		}
	}
	return ways
}

func Solve1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	workflows := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		workflows = append(workflows, line)
	}

	parts := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts = append(parts, line)
	}

	return process(workflows, parts), nil
}

func Solve2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	workflows := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		workflows = append(workflows, line)
	}

	return ways(workflows), nil
}

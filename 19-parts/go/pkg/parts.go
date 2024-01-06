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

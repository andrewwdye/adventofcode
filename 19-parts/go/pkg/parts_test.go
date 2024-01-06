package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkflows(t *testing.T) {
	assert.Equal(t, map[string]Workflow{
		"px": {
			Rules: []Rule{
				{'a', '<', 2006, "qkq"},
				{'m', '>', 2090, "A"},
				{Destination: "rfg"},
			},
		},
	}, getWorkflows([]string{"px{a<2006:qkq,m>2090:A,rfg}"}))
}

func TestGetParts(t *testing.T) {
	assert.Equal(t, []Part{
		{'x': 787, 'm': 2655, 'a': 1222, 's': 2876},
	}, getParts([]string{"{x=787,m=2655,a=1222,s=2876}"}))
}

func TestProcessPart(t *testing.T) {
	workflowLines := `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}`
	workflows := getWorkflows(strings.Split(workflowLines, "\n"))

	assert.True(t, processPart(workflows, Part{'x': 787, 'm': 2655, 'a': 1222, 's': 2876}))
	assert.False(t, processPart(workflows, Part{'x': 1679, 'm': 44, 'a': 2067, 's': 496}))
	assert.True(t, processPart(workflows, Part{'x': 2036, 'm': 264, 'a': 79, 's': 2244}))
	assert.False(t, processPart(workflows, Part{'x': 2461, 'm': 1339, 'a': 466, 's': 291}))
	assert.True(t, processPart(workflows, Part{'x': 2127, 'm': 1623, 'a': 2188, 's': 1013}))
}

func TestSolve1(t *testing.T) {
	input := `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

	assert.Equal(t, 19114, lo.Must(Solve1(strings.NewReader(input))))
}

func TestWays(t *testing.T) {
	assert.Equal(t, 4000*4000*4000*1350, ways([]string{"in{s<1351:A,R}"}))
	assert.Equal(t, 4000*4000*4000*(4000-1350), ways([]string{"in{s<1351:R,A}"}))
}

func TestSolve2(t *testing.T) {
	input := `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

	assert.Equal(t, 167409079868000, lo.Must(Solve2(strings.NewReader(input))))
}

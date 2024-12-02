package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Equal(t, 52, hash("HASH"))
}

func TestSolve1(t *testing.T) {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	assert.Equal(t, 1320, lo.Must(Solve1(strings.NewReader(input))))
}

func TestSolve2(t *testing.T) {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	assert.Equal(t, 145, lo.Must(Solve2(strings.NewReader(input))))
}

func TestProcess(t *testing.T) {
	b := boxes{}
	b.process("rn=1")
	assert.Equal(t, boxes{0: []lens{{"rn", 1}}}, b)
	b.process("cm-")
	assert.Equal(t, boxes{0: []lens{{"rn", 1}}}, b)
	b.process("qp=3")
	b.process("cm=2")
	b.process("qp-")
	b.process("pc=4")
	b.process("ot=9")
	b.process("ab=5")
	b.process("pc-")
	b.process("pc=6")
	b.process("ot=7")
	assert.Equal(t, boxes{0: []lens{{"rn", 1}, {"cm", 2}}, 3: []lens{{"ot", 7}, {"ab", 5}, {"pc", 6}}}, b)
}

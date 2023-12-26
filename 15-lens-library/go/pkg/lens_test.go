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

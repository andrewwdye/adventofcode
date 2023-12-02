package main

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestParseLine2(t *testing.T) {
	assert.Equal(t, 29, lo.Must(parseLine2("two1nine")))
	assert.Equal(t, 83, lo.Must(parseLine2("eightwothree")))
	assert.Equal(t, 13, lo.Must(parseLine2("abcone2threexyz")))
	assert.Equal(t, 24, lo.Must(parseLine2("xtwone3four")))
	assert.Equal(t, 42, lo.Must(parseLine2("4nineeightseven2")))
	assert.Equal(t, 14, lo.Must(parseLine2("zoneight234")))
	assert.Equal(t, 76, lo.Must(parseLine2("7pqrstsixteen")))
}

func TestInput(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			`5ffour295
m9qvkqlgfhtwo3seven4seven
2vdqng1sixzjlkjvq
5twonineeight3onefive
2three2seveneightseven
eightsevenfive3bcptwo
five8six
twonineseven24one3
one8bdxplbtfninefourspqn
nineeight3fiveseven
`,
			532,
		},
		{
			// 35 + 62 + 48 + 42 + 59 + 96 + 42 + 45 + 75 + 13 = 517
			`3dxbsctxgntfivehlcbdzgqtxvqddsjdrjnpgjtxhc
6ftlgzrbfjeightsix5onesevenfourtwoneh
4ninezcpvppbktl35eight
four7five27nine2
5qhbdqjcdtbsevenfivenine
99six
fourfivezhgtbmkhxrj9threehtwonebj
4threethree19threefivetjmcnvpkrfdmhjsnzlv
7sevenvxmtninefvvprtdhkhxpkth5
qbmtvl12fiveclone4three
`,
			517,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, lo.Must(sumCalibrationValues(strings.NewReader(test.input))))
	}
}

package main

import (
	"fmt"
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
	assert.Equal(t, 82, lo.Must(parseLine2("8pztdljxbjjthreenineeightseven7crkdr8eightwocb")))

	assert.Equal(t, 35, lo.Must(parseLine2("3dxbsctxgntfivehlcbdzgqtxvqddsjdrjnpgjtxhc")))
	assert.Equal(t, 61, lo.Must(parseLine2("6ftlgzrbfjeightsix5onesevenfourtwoneh")))
	assert.Equal(t, 48, lo.Must(parseLine2("4ninezcpvppbktl35eight")))
	assert.Equal(t, 42, lo.Must(parseLine2("four7five27nine2")))
	assert.Equal(t, 59, lo.Must(parseLine2("5qhbdqjcdtbsevenfivenine")))
	assert.Equal(t, 96, lo.Must(parseLine2("99six")))
	assert.Equal(t, 41, lo.Must(parseLine2("fourfivezhgtbmkhxrj9threehtwonebj")))
	assert.Equal(t, 45, lo.Must(parseLine2("4threethree19threefivetjmcnvpkrfdmhjsnzlv")))
	assert.Equal(t, 75, lo.Must(parseLine2("7sevenvxmtninefvvprtdhkhxpkth5")))
	assert.Equal(t, 13, lo.Must(parseLine2("qbmtvl12fiveclone4three")))
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
nineeight3fiveseven`,
			532,
		},
		{
			// 35 + 61 + 48 + 41 + 59 + 96 + 41 + 45 + 75 + 13 = 515
			`3dxbsctxgntfivehlcbdzgqtxvqddsjdrjnpgjtxhc
6ftlgzrbfjeightsix5onesevenfourtwoneh
4ninezcpvppbktl35eight
four7five27nine2
5qhbdqjcdtbsevenfivenine
99six
fourfivezhgtbmkhxrj9threehtwonebj
4threethree19threefivetjmcnvpkrfdmhjsnzlv
7sevenvxmtninefvvprtdhkhxpkth5
qbmtvl12fiveclone4three`,
			515,
		},
		{
			// 82 + 93 + 86 + 95 + 79 + 66 + 37 + 92 + 42 + 65 = 737
			`8pztdljxbjjthreenineeightseven7crkdr8eightwocb
9krvttdxf34mrpzzchrgeightthree
8jqlmgseveneightzvxrszfsixf
ninefiveztthreeeight7l5
7vhbjdnldvlfourhpptwo53pqbnzqnine
mxdhsixseven6
threepzlkeight7ppdpqqlv
nine15threeqxlrngntwokhzgh
4hblzjb22
pm6eighteightsnztcmfoureightninefive`,
			737,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, test.output, lo.Must(sumCalibrationValues(strings.NewReader(test.input))))
		})
	}
}

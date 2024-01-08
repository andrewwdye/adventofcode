package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`
		modules := getModules(strings.NewReader(input))

		assert.Equal(t, map[Pulse]int{
			Low:  8,
			High: 4,
		}, round(modules))
	})

	t.Run("2", func(t *testing.T) {
		input := `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`
		modules := getModules(strings.NewReader(input))

		assert.Equal(t, map[Pulse]int{
			Low:  4,
			High: 4,
		}, round(modules))
		assert.Equal(t, map[Pulse]int{
			Low:  4,
			High: 2,
		}, round(modules))
		assert.Equal(t, map[Pulse]int{
			Low:  5,
			High: 3,
		}, round(modules))
		assert.Equal(t, map[Pulse]int{
			Low:  4,
			High: 2,
		}, round(modules))
	})
}

func TestRun(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`
		modules := getModules(strings.NewReader(input))

		assert.Equal(t, 32000000, run(modules, 1000))
	})

	t.Run("2", func(t *testing.T) {
		input := `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`
		modules := getModules(strings.NewReader(input))

		assert.Equal(t, 11687500, run(modules, 1000))
	})
}

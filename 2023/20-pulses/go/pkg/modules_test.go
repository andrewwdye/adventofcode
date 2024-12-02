package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetModules(t *testing.T) {
	input := `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`
	modules := getModules(strings.NewReader(input))
	assert.Len(t, modules, 5)
	assert.Equal(t, &BroadcastModule{BaseModule{line: "broadcaster -> a, b, c", name: "broadcaster", destinations: []string{"a", "b", "c"}}, 0}, modules["broadcaster"])
	assert.Equal(t, &FlipFlopModule{BaseModule{line: "%a -> b", name: "a", destinations: []string{"b"}}, 0}, modules["a"])
	assert.Equal(t, &FlipFlopModule{BaseModule{line: "%b -> c", name: "b", destinations: []string{"c"}}, 0}, modules["b"])
	assert.Equal(t, &FlipFlopModule{BaseModule{line: "%c -> inv", name: "c", destinations: []string{"inv"}}, 0}, modules["c"])
	assert.Equal(t, &ConjunctionModule{BaseModule{line: "&inv -> a", name: "inv", destinations: []string{"a"}}, map[string]Pulse{"c": Low}}, modules["inv"])
}

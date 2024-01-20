package pkg

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Pulse int

const (
	Low  Pulse = 0
	High Pulse = 1
)

func (p Pulse) Opposite() Pulse {
	return 1 - p
}

func (p Pulse) String() string {
	switch p {
	case Low:
		return "low"
	case High:
		return "high"
	}
	panic(fmt.Sprintf("invalid pulse: %d", p))
}

type Module interface {
	Name() string
	Destinations() []string
	Send(in Pulse, from string) *Pulse
}

type BaseModule struct {
	line         string
	name         string
	destinations []string
}

type BroadcastModule struct {
	BaseModule
	value Pulse
}

func (m *BroadcastModule) Name() string {
	return m.name
}

func (m *BroadcastModule) Destinations() []string {
	return m.destinations
}

func (m *BroadcastModule) Send(in Pulse, from string) *Pulse {
	m.value = in
	return &m.value
}

type FlipFlopModule struct {
	BaseModule
	value Pulse
}

func (m *FlipFlopModule) Name() string {
	return m.name
}

func (m *FlipFlopModule) Destinations() []string {
	return m.destinations
}

func (m *FlipFlopModule) Send(in Pulse, from string) *Pulse {
	switch in {
	case Low:
		m.value = m.value.Opposite()
		return &m.value
	case High:
		return nil
	}
	panic(fmt.Sprint("invalid pulse: ", in))
}

type ConjunctionModule struct {
	BaseModule
	recentPulses map[string]Pulse
}

func (m *ConjunctionModule) Name() string {
	return m.name
}

func (m *ConjunctionModule) Destinations() []string {
	return m.destinations
}

func (m *ConjunctionModule) Send(in Pulse, from string) *Pulse {
	m.recentPulses[from] = in
	for _, pulse := range m.recentPulses {
		if pulse == Low {
			v := High
			return &v
		}
	}
	v := Low
	return &v
}

type RxModule struct {
	BaseModule
	Activated bool
}

func (m *RxModule) Name() string {
	return m.name
}

func (m *RxModule) Destinations() []string {
	return m.destinations
}

func (m *RxModule) Send(in Pulse, from string) *Pulse {
	if in == Low {
		m.Activated = true
	}
	return nil
}

func getModules(reader io.Reader) map[string]Module {
	modules := make(map[string]Module)
	re := regexp.MustCompile(`(.+) -> (.+)`)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic(fmt.Sprint("invalid line: ", line))
		}
		destinations := strings.Split(matches[2], ", ")
		var module Module
		switch matches[1][0] {
		case 'b':
			module = &BroadcastModule{
				BaseModule: BaseModule{
					line:         line,
					name:         matches[1],
					destinations: destinations,
				},
			}
		case '%':
			module = &FlipFlopModule{
				BaseModule: BaseModule{
					line:         line,
					name:         matches[1][1:],
					destinations: destinations,
				},
			}
		case '&':
			module = &ConjunctionModule{
				BaseModule: BaseModule{
					line:         line,
					name:         matches[1][1:],
					destinations: destinations,
				},
				recentPulses: make(map[string]Pulse),
			}
		}
		modules[module.Name()] = module
	}
	// Add an rx module
	modules["rx"] = &RxModule{
		BaseModule: BaseModule{
			name: "rx",
		},
		Activated: false,
	}
	for _, module := range modules {
		for _, dest := range module.Destinations() {
			if destModule, ok := modules[dest]; ok {
				if con, ok := destModule.(*ConjunctionModule); ok {
					con.recentPulses[module.Name()] = Low
				}
			}
		}
	}
	return modules
}

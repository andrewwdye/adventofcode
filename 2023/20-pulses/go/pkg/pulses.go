package pkg

import (
	"fmt"
	"io"
)

type Action struct {
	to    Module
	from  string
	pulse Pulse
}

func round(modules map[string]Module) map[Pulse]int {
	start := modules["broadcaster"]
	actions := []Action{{start, "button", Low}}
	pulses := make(map[Pulse]int, 2)
	for len(actions) > 0 {
		action := actions[0]
		actions = actions[1:]
		pulses[action.pulse] += 1
		if action.to == nil {
			continue
		}
		result := action.to.Send(action.pulse, action.from)
		if result != nil {
			for _, dest := range action.to.Destinations() {
				next := modules[dest]
				// fmt.Printf("%s -%s-> %s\n", action.to.Name(), *result, dest)
				actions = append(actions, Action{next, action.to.Name(), *result})
			}
		}
	}
	return pulses
}

func run(modules map[string]Module, count int) int {
	totals := make(map[Pulse]int, 2)
	for i := 0; i < count; i++ {
		pulses := round(modules)
		for pulse, count := range pulses {
			totals[pulse] += count
		}
	}
	return totals[Low] * totals[High]
}

func run2(modules map[string]Module) int {
	count := 0
	for {
		round(modules)
		count += 1
		if vd, ok := modules["vd"]; ok {
			vd := vd.(*ConjunctionModule)
			for source, p := range vd.recentPulses {
				if p == High {
					fmt.Printf("source: %s, count: %d\n", source, count)
				}
			}
		}
		if rx, ok := modules["rx"]; ok && rx.(*RxModule).Activated {
			return count
		}
	}
}

func Solve1(reader io.Reader) (int, error) {
	modules := getModules(reader)
	// for name, m := range modules {
	// 	fmt.Println(name, m)
	// }
	return run(modules, 1000), nil
}

func Solve2(reader io.Reader) (int, error) {
	modules := getModules(reader)
	// for name, m := range modules {
	// 	fmt.Println(name, m)
	// }
	// g := NewGraph(modules)
	// f, err := os.Create("/tmp/graph.dot")
	// if err != nil {
	// 	return 0, err
	// }
	// g.OutputDOT(f)
	return run2(modules), nil
}

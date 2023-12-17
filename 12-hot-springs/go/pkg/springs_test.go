package pkg

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCountWays(t *testing.T) {
	t.Run("all known", func(t *testing.T) {
		lines := []string{
			"#.#.### 1,1,3",
			".#...#....###. 1,1,3",
			".#.###.#.###### 1,3,1,6",
			"####.#...#... 4,1,1",
			"#....######..#####. 1,6,5",
			".###.##....# 3,2,1",
		}
		for _, line := range lines {
			assert.Equal(t, 1, lo.Must(countWays(line, false)), line)
		}
	})

	t.Run("unknown", func(t *testing.T) {
		assert.Equal(t, 1, lo.Must(countWays("???.### 1,1,3", false)))
		assert.Equal(t, 4, lo.Must(countWays(".??..??...?##. 1,1,3", false)))
		assert.Equal(t, 1, lo.Must(countWays("?#?#?#?#?#?#?#? 1,3,1,6", false)))
		assert.Equal(t, 1, lo.Must(countWays("????.#...#... 4,1,1", false)))
		assert.Equal(t, 4, lo.Must(countWays("????.######..#####. 1,6,5", false)))
		assert.Equal(t, 10, lo.Must(countWays("?###???????? 3,2,1", false)))
	})
}

func TestCountWaysExpanded(t *testing.T) {
	t.Run("all known", func(t *testing.T) {
		lines := []string{
			"#.#.### 1,1,3",
			".#...#....###. 1,1,3",
			".#.###.#.###### 1,3,1,6",
			"####.#...#... 4,1,1",
			"#....######..#####. 1,6,5",
			".###.##....# 3,2,1",
		}
		for _, line := range lines {
			assert.Equal(t, 1, lo.Must(countWays(line, true)), line)
		}
	})

	t.Run("unknown", func(t *testing.T) {
		assert.Equal(t, 1, lo.Must(countWays("???.### 1,1,3", true)))
		assert.Equal(t, 16384, lo.Must(countWays(".??..??...?##. 1,1,3", true)))
		assert.Equal(t, 1, lo.Must(countWays("?#?#?#?#?#?#?#? 1,3,1,6", true)))
		assert.Equal(t, 16, lo.Must(countWays("????.#...#... 4,1,1", true)))
		assert.Equal(t, 2500, lo.Must(countWays("????.######..#####. 1,6,5", true)))
		assert.Equal(t, 506250, lo.Must(countWays("?###???????? 3,2,1", true)))
	})

	t.Run("long", func(t *testing.T) {
		line := "???????#??#.???????????????#??#.???????????????#??#.???????????????#??#.???????????????#??#.??????? 3,2,2,1,2,3,2,2,1,2,3,2,2,1,2,3,2,2,1,2,3,2,2,1,2"
		assert.Equal(t, 1, lo.Must(countWays(line, true)))
	})
}

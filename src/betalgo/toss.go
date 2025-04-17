package betalgo

import (
	"math/rand"
)

type Toss struct{}

// TODO:  Make Sure That there is no more than 5 concurrent or next to next result
// just make it a bit difficult to get a truthy value but it is even if weak
func (t *Toss) Basic(chance, out uint32) int {
	random := uint32(rand.Intn(int(out - 1)))
	if random >= chance {
		return 1
	}
	return 0
}

// the difficulty will reduce or increase every draw
func (t *Toss) DynamicDiffuculty(chance, out uint32, incr int) func() int {
	drawCount := 0
	currentDifficulty := 0

	draw := func() int {
		nchance := chance + uint32(currentDifficulty)
		drawCount++
		currentDifficulty = (drawCount * int(incr))
		return t.Basic(nchance, out)

	}
	return draw
}

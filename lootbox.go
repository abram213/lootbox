package lootbox

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

type LootBox struct {
	weights []int
	props   props
}

type props struct {
	limits []int
	sum    int
}

func New(weights []int) (*LootBox, error) {
	if len(weights) == 0 {
		return nil, errors.New("weights length should be > 0")
	}

	var sum int
	var limits []int
	for _, w := range weights {
		limits = append(limits, w+sum)
		sum += w
	}
	return &LootBox{weights, props{limits, sum}}, nil
}

func (lb *LootBox) Randomise() int {
	rand.Seed(time.Now().UnixNano())
	return sort.SearchInts(lb.props.limits, rand.Intn(lb.props.sum-1)+1)
}

func (lb *LootBox) DropChances() []float64 {
	var chances []float64
	for _, w := range lb.weights {
		chances = append(chances, (float64(w)/float64(lb.props.sum))*100)
	}
	sort.Float64s(chances)
	return chances
}

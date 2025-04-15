package betalgo

import (
	"errors"
	"slices"
)

// The betting algorithim to calculate players returns
type Facet struct {
	ID         int
	name       string
	Likelyhood int
}

// This is used to create a new posible outcome
func NewFacet(id int, name string, lk int) *Facet {
	return &Facet{ID: id, name: name, Likelyhood: lk}
}

func (c *Facet) positions() []int {
	ids := make([]int, c.Likelyhood)
	for i := range c.Likelyhood {
		ids[i] = c.ID
	}
	return ids
}

func InitFacets(choices []Facet, target int) ([]int, error) {
	total := int(0)
	options := []int{}
	for _, choice := range choices {
		if choice.Likelyhood < 0 {
			return nil, errors.New("Betalgo.InitChoices: Invalid likelyhood ")
		} else {
			total += choice.Likelyhood
		}
	}
	if total != target {
		return nil, errors.New("Betalgo.InitChoices: Invalid Total Target")
	}

	for _, choice := range choices {
		nums := choice.positions()
		options = append(options, nums...)
	}
	return options, nil
}

func SpreadEven(values []int) []int {
	skip := []int{}
	data := slices.Clone(skip)
	prev := 0
	for _, value := range values {
		if prev == value {
			skip = append(skip, value)
		} else {
			data = append(data, value)
		}
		prev = value
	}

	if len(skip) > 0 {
		appended := append(data, skip...)
		return SpreadEven(appended)
	}
	return data
}

// TODO: Shuffle

type ShuffleOptions struct {
}

// TODO: Dealings
type DealData struct {
	// the list of users
	Users int
	// The items will be shown as [user][cards]
	Items [][]int
}

package betalgo

import (
	"errors"
	"log"
	"math/rand"
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

func SpreadEven(choices []Facet, target int) []int {
	current := 0
	data := make([]int, 0)
	empty := 0
	for i := 0; i < target; {

		if empty == len(choices) {
			log.Fatalf("betalgo.SpreadEven: Invalid Likelyhood distribution")
		}
		if current == len(choices) {
			current = 0
		}
		choice := &choices[current]
		if choice.Likelyhood <= 0 {
			empty++
		} else {
			data = append(data, choice.ID)
			empty = 0
			choice.Likelyhood--
			i++
		}

		current++
	}
	return data
}

// TODO: Shuffle
type Shuff struct {
	Items []int
}

func (c *Shuff) Shuffle() {
	length := len(c.Items)
	for i, v := range c.Items {
		random := rand.Intn(length - 1)
		cpy := c.Items[random]
		c.Items[i] = cpy
		c.Items[random] = v
	}
}

// TODO: Dealings
type DealerData struct {
	// the list of users
	Users int
	// The items will be shown as [user][cards]
	Items [][]int
}

func Deal(items []int, users, todistribute int) *DealerData {
	data := new(DealerData)
	data.Items = make([][]int, users)
	data.Users = users
	current := 0
	for i := range todistribute {
		// reset every users
		if current == users {
			current = 0
		}
		data.Items[current] = append(data.Items[current], items[i])
		current++
	}

	return data
}

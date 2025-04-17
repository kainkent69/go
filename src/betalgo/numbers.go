package betalgo

import (
	"errors"
	"math/rand"
	"slices"
)

// The options for making random Number
type RandomNumbersOptions struct {
	// Dect is the list of numbers
	Dect []int32
	// if the numbers are allowed to be repeated in every number
	Reapeats bool
	// Any order are allowed to be done
	AnyOrder bool
	// The number of times of draw
	DrawCount int
}

func NewRandomOptions(dect []int32, repeat,
	anyOrder bool, count int) *RandomNumbersOptions {
	return &RandomNumbersOptions{Dect: dect, Reapeats: repeat, AnyOrder: anyOrder, DrawCount: count}
}

func (opt RandomNumbersOptions) GenerateRandom() []int {
	draw := func() int {
		items := len(opt.Dect)
		random := rand.Intn(items)
		opt.DrawCount--
		if !opt.Reapeats {
			opt.Dect = slices.Delete(opt.Dect, random-1, random)
		}
		return random
	}

	result := []int{}
	for opt.DrawCount > 0 {
		result = append(result, draw())
	}

	return result
}

// / CompareNumbers check if the bumbers are equal based on the conditions of the generition
func (opt RandomNumbersOptions) CompareNumbers(choice, result []int) (bool, error) {
	if len(choice) != len(choice) {
		return false, errors.New("betalgo.CompareNumbers: does not have a matching length")
	}
	if opt.AnyOrder {
		cpy := slices.Clone(choice)
		cpy2 := slices.Clone(result)
		slices.Sort(cpy)
		slices.Sort(cpy2)
		return slices.Equal(cpy, cpy2), nil
	} else {
		return slices.Equal(choice, result), nil
	}

}

// Generate random and compare numberes
func (opt RandomNumbersOptions) GenerateCompareNumbers(choice []int) (bool, error) {
	random := opt.GenerateRandom()
	return opt.CompareNumbers(choice, random)
}

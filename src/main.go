package main

import (
	"bingof/betalgo"
	"fmt"
	"log"
)

func main() {
	shuf := new(betalgo.Shuff)
	shuf.Items = make([]int, 0)
	for i := range 100 {
		shuf.Items = append(shuf.Items, i+1)
	}

	log.Printf("Before the shuffle %v\n", shuf)
	shuf.Shuffle()
	log.Printf("Shuffled %v\n", shuf)
	log.Printf("Shuffled lenth %v\n", len(shuf.Items))
	fmt.Print("\n\n\n\n")
	deal := betalgo.Deal(shuf.Items, 4, 49)
	fmt.Printf("DEAL\n%v\n", deal)
	fmt.Printf("\n\n\n\n")

	facet0 := betalgo.NewFacet(1, "Facet0", 30)
	facet1 := betalgo.NewFacet(2, "Facet1", 30)
	facet2 := betalgo.NewFacet(3, "Facet2", 30)
	facet3 := betalgo.NewFacet(4, "Facet3", 10)

	// initialize facet
	ids, err := betalgo.InitFacets([]betalgo.Facet{*facet0, *facet1, *facet2, *facet3}, 100)
	if err != nil {
		log.Panic(err)
	}
	_ = ids

	even := betalgo.SpreadEven([]betalgo.Facet{*facet0, *facet1, *facet2, *facet3}, 100)
	log.Print(even)

	fmt.Print("\n\n\n\n")
	log.Println("LOGGER=======================TOSS=============================REGGOL")
	for range 10 {
		toss := new(betalgo.Toss)
		res := toss.Basic(50, 100)
		log.Printf("Toss Result: %d\n", res)
	}

	fmt.Print("\n\n\n\n")
	number := *betalgo.NewRandomOptions([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true, false, 3)

	for range 10 {

		random := number.GenerateRandom()
		log.Printf("Number Result In Exact: %d\n", random)
		res, err := number.GenerateCompareNumbers(random)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Did you win? %t\n", res)

	}
}

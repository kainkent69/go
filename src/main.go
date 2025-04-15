package main

import (
	"bingof/betalgo"
	"log"
)

func main() {
	// choice1 := betalgo.NewChoice(1, "choioce1", 50)
	// choice2 := betalgo.NewChoice(2, "choioce", 50)
	// choices := []betalgo.Choice{*choice1, *choice2}
	//
	// options, err := betalgo.CreateChoices(choices, 100)
	// if err != nil {
	// 	log.Fatalf("Test: Something When wrong %s", err)
	// }
	//
	// log.Printf("values: %v", options)
	//
	choice1 := betalgo.NewChoice(1, "choioce1", 25)
	choice2 := betalgo.NewChoice(2, "choioce2", 25)

	choice3 := betalgo.NewChoice(3, "choioce3", 20)
	choice4 := betalgo.NewChoice(4, "choioce4", 10)
	choice5 := betalgo.NewChoice(5, "choioce4", 20)

	choices := []betalgo.Facetj{*choice1, *choice2, *choice3, *choice4, *choice5}

	options, err := betalgo.SpreadEven(choices, 100)
	if err != nil {
		log.Fatalf("Test: Something When wrong %s", err)
	}
	log.Printf("\n\nvalues: %v", options)

	random := betalgo.SpreadEven(options)
	options2 := []int{1, 2, 3, 3, 3, 3, 4, 4, 5, 6, 6, 5, 5, 5, 5}

	log.Printf("\n\n Options randomized is this %v\n", random)
	log.Printf("\n\n Options randomized2 is this %v\n", betalgo.SpreadEven(options2))

}

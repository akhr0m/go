package main

import (
	"chapter-71-exercise-ninja-level-13/test2/quote"
	"chapter-71-exercise-ninja-level-13/test2/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.AlberEinsteinSaid))           //30            //we can see how much words on quotes of Albert einstein
	for i, v := range word.UseCount(quote.AlberEinsteinSaid) { //printed array //we can print every word that in the quotes
		fmt.Println(v, i)
	}

	d := word.Count("andi subagja dirjo mangunharjo sujatmiko lamunglaksono")
	fmt.Println(d, "---> word(s)")
}
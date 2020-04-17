package lootbox

import "fmt"

type item struct {
	name   string
	weight int
}

func gameLootBox() {
	var items []item = []item{
		{"Sword", 10},
		{"Knife", 20},
		{"Shield", 20},

		{"Rare Sword", 5},
		{"Rare Knife", 10},
		{"Rare Shield", 10},

		{"Epic Sword", 3},
		{"Epic Knife", 5},
		{"Epic Shield", 5},

		{"Nothing", 100},
	}

	//prepare the data
	var weights []int
	for _, item := range items {
		weights = append(weights, item.weight)
	}

	//create new loot box
	lootbox, err := New(weights)
	if err != nil {
		fmt.Println(err)
	}

	//get the chances of the item dropping out
	chances := lootbox.DropChances()
	for i, item := range items {
		fmt.Println(item.name, " - ", chances[i], "%")
	}

	//randomise item drop 10 times
	for i := 0; i < 10; i++ {
		fmt.Print("Your get: ", items[lootbox.Randomise()].name, "\n")
	}
}

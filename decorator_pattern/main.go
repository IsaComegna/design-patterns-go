package main

import (
	"fmt"

	beverage "github.com/IsaComegna/design-patterns-go/beverage"
	decorator "github.com/IsaComegna/design-patterns-go/decorator"
)

func main() {
	espresso := beverage.NewEspresso()
	latte := decorator.NewBeverageWithMilk(espresso)
	latteWithChocolate := decorator.NewBeverageWithChocolate(latte)

	fmt.Println("1) ", espresso.GetFullDescription())
	fmt.Println("2) ", latte.GetFullDescription())
	fmt.Println("3) ", latteWithChocolate.GetFullDescription())
}

package decorator

import (
	"fmt"

	beverage "github.com/IsaComegna/design-patterns-go/beverage"
)

const chocolatePrice float32 = 3.00
const chocolateDescription string = "chocolate"

type BeverageWithChocolate struct {
	BaseBeverage beverage.Beverage
}

func NewBeverageWithChocolate(beverage beverage.Beverage) BeverageWithChocolate {
	return BeverageWithChocolate{
		BaseBeverage: beverage,
	}
}

func (b BeverageWithChocolate) GetCost() float32 {
	return chocolatePrice + b.BaseBeverage.GetCost()
}

func (b BeverageWithChocolate) GetDescription() string {
	return fmt.Sprintf("%s with %s", b.BaseBeverage.GetDescription(), chocolateDescription)
}

func (b BeverageWithChocolate) GetFullDescription() string {
	return fmt.Sprintf("Order (%s) total cost = %v", b.GetDescription(), b.GetCost())
}

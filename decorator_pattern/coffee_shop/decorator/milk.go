package decorator

import (
	"fmt"

	beverage "github.com/IsaComegna/design-patterns-go/coffee_shop/beverage"
)

const milkPrice float32 = 5.50
const milkDescription string = "milk"

type BeverageWithMilk struct {
	BaseBeverage beverage.Beverage
}

func NewBeverageWithMilk(beverage beverage.Beverage) BeverageWithMilk {
	return BeverageWithMilk{
		BaseBeverage: beverage,
	}
}

func (b BeverageWithMilk) GetCost() float32 {
	return milkPrice + b.BaseBeverage.GetCost()
}

func (b BeverageWithMilk) GetDescription() string {
	return fmt.Sprintf("%s with %s", b.BaseBeverage.GetDescription(), milkDescription)
}

func (b BeverageWithMilk) GetFullDescription() string {
	return fmt.Sprintf("Order (%s) total cost = %v", b.GetDescription(), b.GetCost())
}

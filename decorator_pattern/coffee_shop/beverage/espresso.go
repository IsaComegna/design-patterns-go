package beverage

import "fmt"

const espressoPrice float32 = 12.90
const espressoDescription string = "espresso"

type Espresso struct{}

func NewEspresso() Espresso {
	return Espresso{}
}

func (e Espresso) GetCost() float32 {
	return espressoPrice
}

func (e Espresso) GetDescription() string {
	return espressoDescription
}

func (e Espresso) GetFullDescription() string {
	return fmt.Sprintf("Order (%s) total cost = %v", e.GetDescription(), e.GetCost())
}

package main

import (
	"fmt"

	"github.com/IsaComegna/design-patterns-go/travel_costs_calculator/strategy"
)

func main() {
	var travelTheWorld []strategy.Travel

	busTravel := strategy.NewBusTravel(0)
	airplaneTravel := strategy.NewAirplaneTravel(1)

	travelTheWorld = append(travelTheWorld, busTravel)

	fmt.Println("Price bus", calculateTotalPrice(travelTheWorld))

	travelTheWorld = append(travelTheWorld, airplaneTravel)

	fmt.Println("Price bus + airplane", calculateTotalPrice(travelTheWorld))
}

func calculateTotalPrice(travelArray []strategy.Travel) float32 {
	var totalPrice float32 = 0

	for _, travel := range travelArray {
		totalPrice = totalPrice + travel.CalculatePrice()
	}

	return totalPrice
}

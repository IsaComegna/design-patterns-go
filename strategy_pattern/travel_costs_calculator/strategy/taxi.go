package strategy

type TaxiTravel struct {
	Distance float32
}

const taxiKilometerPrice float32 = 1.15

func NewTaxiTravel(distance float32) TaxiTravel {
	return TaxiTravel{
		Distance: distance,
	}
}

func (tt TaxiTravel) CalculatePrice() float32 {
	return tt.Distance * taxiKilometerPrice
}

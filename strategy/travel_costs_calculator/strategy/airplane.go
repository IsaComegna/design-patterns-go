package strategy

type AirplaneTravel struct {
	NumberOfTransfers int
}

const airplaneTicketPrice int = 52000

func NewAirplaneTravel(numberOfTransfers int) AirplaneTravel {
	return AirplaneTravel{
		NumberOfTransfers: numberOfTransfers,
	}
}

func (at AirplaneTravel) CalculatePrice() float32 {
	return float32((airplaneTicketPrice * (1 + at.NumberOfTransfers)) / 100)
}

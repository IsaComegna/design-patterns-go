package strategy

type BusTravel struct {
	NumberOfTransfers int
}

const busTicketPrice int = 345

func NewBusTravel(numberOfTransfers int) BusTravel {
	return BusTravel{
		NumberOfTransfers: numberOfTransfers,
	}
}

func (bt BusTravel) CalculatePrice() float32 {
	return float32((busTicketPrice * (1 + bt.NumberOfTransfers)) / 100)
}

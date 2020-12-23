package beverage

type Beverage interface {
	GetCost() float32
	GetDescription() string
	GetFullDescription() string
}

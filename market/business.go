package market

type BusType int

const(
	Landscaping BusType = iota
	Pizza
	Barber
)

type Business struct{
	BusinessType BusType
	BusinessName string
	Owner *User
}
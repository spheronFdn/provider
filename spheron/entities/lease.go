package entities

type Lease struct {
	OrderID         uint64
	ProviderAddress string
	AcceptedPrice   uint64
	State           OrderState
}

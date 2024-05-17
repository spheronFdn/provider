package entities

type Lease struct {
	OrderID       uint64
	Creator       string
	Provider      string
	AcceptedPrice uint64
	State         OrderState
}

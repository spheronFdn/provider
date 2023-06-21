package bidengine

// Status stores orders
type Status struct {
	Orders uint32 `json:"orders"`
}

type DepositorsData struct {
	Depositors []string `json:"depositors"`
}

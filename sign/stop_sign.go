package sign

// StopSign interface .
type StopSign interface {

	// Sign send
	Sign() bool

	// Signed return
	Signed() bool

	// Reset Sign
	Reset()

	// Deal Sign
	Deal(code string)

	// DealCount
	DealCount(code string) uint32

	// DealTotal
	DealTotal() uint32

	// Summary
	Summary() string
}

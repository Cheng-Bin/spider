package middleware

// IDGenerator interface.
type IDGenerator interface {
	GetUint32() uint32
}

package base

// Item type
type Item map[string]interface{}

// Valid return item valid .
// Data interface impl .
func (item Item) Valid() bool {
	return item != nil
}

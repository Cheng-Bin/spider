package itempipeline

import (
	"github.com/chapin/spider/base"
)

// ProcessItem type.
type ProcessItem func(item base.Item) (result base.Item, err error)

// ItemPipeline interface.
type ItemPipeline interface {
	// Send item
	Send(item base.Item) []error

	// FailFast return bool.
	FailFast() bool

	// SetFailFast flag.
	SetFailFast(failFast bool)

	// Count present finished task.
	Count() []uint64

	// ProcessingNumber task.
	ProcessingNumber() uint64

	// Summary info.
	Summary() string
}

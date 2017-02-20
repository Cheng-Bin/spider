package middleware

import (
	"math"
	"sync"
)

type cyclicIDGenertor struct {
	sn    uint32
	ended bool
	mutex sync.Mutex
}

func NewIDGenertor() IDGenerator {
	return &cyclicIDGenertor{}
}

func (gen *cyclicIDGenertor) GetUint32() uint32 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.ended {
		defer func() { gen.ended = false }()
		gen.sn = 0
		return gen.sn
	}
	id := gen.sn
	if id < math.MaxUint32 {
		gen.sn++
	} else {
		gen.ended = true
	}

	return id
}

type cyclicIdGenertor2 struct {
	base       cyclicIDGenertor
	cycleCount int64
}

func (gen *cyclicIdGenertor2) GetUint64() uint64 {
	var id64 uint64
	if gen.cycleCount%2 == 1 {
		id64 += math.MaxUint32
	}
	id32 := gen.base.GetUint32()
	if id32 == math.MaxUint32 {
		gen.cycleCount++
	}
	id64 += uint64(id32)
	return id64
}

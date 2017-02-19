package pool

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type myPool struct {
	total       uint32        // pool capcity
	etype       reflect.Type  // Entity type
	genEntity   func() Entity // Entity gen
	container   chan Entity   // Entity
	idContainer map[uint32]bool
	mutex       sync.Mutex
}

func NewPool(total uint32,
	entityType reflect.Type,
	genEntity func() Entity) (Pool, error) {

	if total == 0 {
		errMsg := fmt.Sprintf("The pool can not be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}

	size := int(total)
	container := make(chan Entity, size)
	idContainer := make(map[uint32]bool)
	for i := 0; i < size; i++ {
		newEntity := genEntity()
		if entityType != reflect.TypeOf(newEntity) {
			errMsg := fmt.Sprintf("The type of result of function genEntity() is Not %s!\n", entityType)
			return nil, errors.New(errMsg)
		}
		container <- newEntity
		idContainer[newEntity.Id()] = true
	}

	pool := &myPool{
		total:       total,
		etype:       entityType,
		genEntity:   genEntity,
		container:   container,
		idContainer: idContainer,
	}

	return pool, nil
}

func (pool *myPool) Take() (Entity, error) {
	entity, ok := <-pool.container
	if !ok {
		return nil, errors.New("The inner container is invalid!")
	}

	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	pool.idContainer[entity.Id()] = false
	return entity, nil
}

func (pool *myPool) Return(entity Entity) error {
	if entity == nil {
		return errors.New("The returning entity is invalid!")
	}

	if pool.etype != reflect.TypeOf(entity) {
		errMsg := fmt.Sprintf("The type of returning entity is NOT %s!\n", pool.etype)
		return errors.New(errMsg)
	}

	entityId := entity.Id()
	casResult := pool.compareAndSetForIdContainer(entityId, false, true)

	if casResult == 1 {
		pool.container <- entity
		return nil
	} else if casResult == 0 {
		errMsg := fmt.Sprintf("The entity (id=%d) is already in the pool!\n", entity.Id())
		return errors.New(errMsg)
	} else {
		errMsg := fmt.Sprintf("The entity (id = %d) is illegal\n", entity.Id())
		return errors.New(errMsg)
	}
}

func (pool *myPool) compareAndSetForIdContainer(entityId uint32,
	oldValue bool, newValue bool) int8 {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	v, ok := pool.idContainer[entityId]
	if !ok {
		return -1
	}
	if v != oldValue {
		return 0
	}

	pool.idContainer[entityId] = newValue
	return 1
}

func (pool *myPool) Total() uint32 {
	return pool.total
}

func (pool *myPool) Used() uint32 {
	return pool.total - uint32(len(pool.idContainer))
}

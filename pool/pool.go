package pool

type Pool interface {
	Take() (Entity, error)
	Return(entity Entity) error
	Total() uint32
	Used() uint32
}

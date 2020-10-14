package repos

import (
	"github.com/jackc/pgx/pgxpool"
)

//EntityProvider describes GetEntity method to get entity data
//we need it to implement as Entity to provide repositories throught
//interfaces
type EntityProvider interface {
	getEntity() *Entity
}

//Entity describes basic entity data structure
type Entity struct {
	ID int
}

//getEntity returns entity data
func (e *Entity) getEntity() *Entity {
	return e
}

//Repository interface describes methods for basic repositories
type Repository interface {
	GetSlice(db *pgxpool.Pool, from int, to int) ([]*EntityProvider, error)
}

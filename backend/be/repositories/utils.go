package repos

import (
	"github.com/jackc/pgx/pgxpool"

	utils "../utils"
)

//IncorrectID is used to mark empty record
const IncorrectID int = -1

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
	SetDb(db *pgxpool.Pool)
	GetSlice(from int, to int) ([]*EntityProvider, error)
	GetEntityByID(id int) (*EntityProvider, error)
	RemoveEntityByID(id int) (int, error)
}

//EntityRepository describes basic repository structure
type EntityRepository struct {
	db *pgxpool.Pool
	log *utils.Logger
	token *utils.TokenGenerator

	Repository
}

//SetDb injects db connection into repository
func (r *EntityRepository) SetDb(db *pgxpool.Pool) {
	r.db = db
}

//SetLogger injects logger into repository
func (r *EntityRepository) SetLogger(logger *utils.Logger) {
	r.log = logger
}

//SetTokenGenerator injects logger into repository
func (r *EntityRepository) SetTokenGenerator(token *utils.TokenGenerator) {
	r.token = token
}

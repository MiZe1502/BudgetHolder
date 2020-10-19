package repos

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

//Shop represents shop data item
type Shop struct {
	Name    string
	Address string
	Comment string
	URL     string

	Entity
}

//ShopsRepository is a repository for shops data
type ShopsRepository struct {
	EntityRepository
}

//GetSlice returns slice of shops data
func (r *ShopsRepository) GetSlice(from int, to int) ([]*Shop, error) {
	var shops []*Shop

	err := pgxscan.Select(context.Background(), r.db, &shops, `SELECT * from budget.get_shops($1, $2)`, from, to)
	if err != nil {
		return nil, err
	}

	return shops, err
}

//GetEntityByID returns shop`s data
func (r *ShopsRepository) GetEntityByID(id int) (Shop, error) {
	var shop Shop

	err := pgxscan.Get(context.Background(), r.db, &shop, `SELECT * from budget.get_shop_by_id($1)`, id)
	if err != nil {
		return shop, err
	}

	return shop, err
}

//RemoveEntityByID marks single shop as removed in db
func (r *ShopsRepository) RemoveEntityByID(id int, uuid string) (int, error) {
	var removedShopID int

	err := pgxscan.Get(context.Background(), r.db, &removedShopID, `SELECT * from budget.remove_shop($1, $2)`, id, uuid)
	if err != nil {
		return IncorrectID, err
	}

	return removedShopID, err
}

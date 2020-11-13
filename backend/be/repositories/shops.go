package repos

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

//Shop represents shop data item
type Shop struct {
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Comment string `json:"comment,omitempty"`
	URL     string `json:"url,omitempty"`

	Entity
}

//SimpleShop represents shop data item for dropdown menus
type SimpleShop struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

//ShopsRepository is a repository for shops data
type ShopsRepository struct {
	EntityRepository
}

//UpdateShop updates existing shop and returns its id
func (r *ShopsRepository) UpdateShop(shopData *Shop, sessionUUID string) (int, error) {
	var updatedShopID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&updatedShopID,
		`SELECT * from budget.update_shop($1, $2, $3, $4, $5, $6::uuid)`,
		shopData.ID,
		shopData.Name,
		shopData.URL,
		shopData.Address,
		shopData.Comment,
		sessionUUID)
	if err != nil {
		return IncorrectID, err
	}

	return updatedShopID, err
}

//GetTopShopsByName finds top shop items by name
func (r *ShopsRepository) GetTopShopsByName(name string) ([]*SimpleShop, error) {
	var shops []*SimpleShop

	err := pgxscan.Select(context.Background(), r.db, &shops, `SELECT * from budget.get_top_shops_by_name($1, $2)`, name, TopCounter)
	if err != nil {
		return nil, err
	}

	return shops, err
}

//CreateNewShop creates new shop and returns its id
func (r *ShopsRepository) CreateNewShop(shopData *Shop, sessionUUID string) (int, error) {
	var addedShopID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&addedShopID,
		`SELECT * from budget.create_new_shop($1, $2, $3, $4, $5::uuid)`,
		shopData.Name,
		shopData.URL,
		shopData.Address,
		shopData.Comment,
		sessionUUID)
	if err != nil {
		return IncorrectID, err
	}

	return addedShopID, err
}

//GetSlice returns slice of shops data
func (r *ShopsRepository) GetSlice(from int, count int) ([]*Shop, error) {
	var shops []*Shop

	err := pgxscan.Select(context.Background(), r.db, &shops, `SELECT * from budget.get_shops($1, $2)`, from, count)
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

package repos

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

//Purchase represents purchase basic data
type Purchase struct {
	TotalPrice float32 `json:"total_price"`
	ShopID     int     `json:"shop_id"`
	Date       int32   `json:"date"`
	Comment    string  `json:"comment"`

	Entity
}

//PurchasesRepository is a repository for purchases data
type PurchasesRepository struct {
	EntityRepository
}

//GetEntityByID returns purchases`s data
func (r *PurchasesRepository) GetEntityByID(id int) (Purchase, error) {
	var purchase Purchase

	err := pgxscan.Get(context.Background(), r.db, &purchase, `SELECT * from budget.get_purchase_by_id($1)`, id)
	if err != nil {
		return purchase, err
	}

	return purchase, err
}

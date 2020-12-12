package repos

import (
	"context"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

//Purchase represents purchase basic data
type Purchase struct {
	TotalPrice float32 `json:"total_price,omitempty"`
	ShopID     int     `json:"shop_id,omitempty"`
	Date       time.Time     `json:"date,omitempty"`
	Comment    string  `json:"comment,omitempty"`

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

//GetSlice returns slice of purchases data
func (r *PurchasesRepository) GetSlice(from int, count int, userID int) ([]*Purchase, error) {
	var purchases []*Purchase

	err := pgxscan.Select(context.Background(), r.db, &purchases, `SELECT * from budget.get_purchases($1, $2, $3)`, from, count, userID)
	if err != nil {
		return nil, err
	}

	return purchases, err
}

//RemoveEntityByID marks single purchase as removed in db
func (r *PurchasesRepository) RemoveEntityByID(id int, uuid uuid.UUID) (int, error) {
	var removedPurchaseID int

	err := pgxscan.Get(context.Background(), r.db, &removedPurchaseID, `SELECT * from budget.remove_purchase($1, $2::uuid)`, id, uuid)
	if err != nil {
		return IncorrectID, err
	}

	return removedPurchaseID, err
}

//CreateNewPurchase creates new purchase and returns its id
func (r *PurchasesRepository) CreateNewPurchase(purchaseData *Purchase, sessionUUID uuid.UUID) (int, error) {
	var addedPurchaseID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&addedPurchaseID,
		`SELECT * from budget.create_new_purchase($1, $2, $3, $4, $5::uuid)`,
		purchaseData.TotalPrice,
		purchaseData.ShopID,
		purchaseData.Date,
		purchaseData.Comment,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return addedPurchaseID, err
}

//UpdatePurchase updates existing purchase and returns its id
func (r *PurchasesRepository) UpdatePurchase(purchaseData *Purchase, sessionUUID uuid.UUID) (int, error) {
	var updatedPurchaseID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&updatedPurchaseID,
		`SELECT * from budget.update_purchase($1, $2, $3, $4, $5, $6::uuid)`,
		purchaseData.ID,
		purchaseData.TotalPrice,
		purchaseData.ShopID,
		purchaseData.Date,
		purchaseData.Comment,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return updatedPurchaseID, err
}

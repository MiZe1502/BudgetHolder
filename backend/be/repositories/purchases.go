package repos

import (
	"context"
	"errors"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

//Purchase represents purchase basic data
type Purchase struct {
	TotalPrice float32 `json:"total_price,omitempty"`
	ShopID     *int    `json:"shop_id,omitempty"`
	Shop	   *Shop   `json:"shop,omitempty"`
	Date       *time.Time   `json:"date,omitempty"`
	Comment    string  `json:"comment,omitempty"`

	Entity
}

//PurchaseWithGoods represents purchase data with corresponding goods
type PurchaseWithGoods struct {
	Purchase

	GoodsData []*GoodsItemWithDetails `json:"goods_data,omitempty"`
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
func (r *PurchasesRepository) GetSlice(from int, count int, userID int) ([]*PurchaseWithGoods, error) {
	var purchases []*PurchaseWithGoods

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

//CountTotal returns total purchases items count
func (r *PurchasesRepository) CountTotal() (int, error) {
	var total int

	err := pgxscan.Get(context.Background(),
		r.db,
		&total,
		`SELECT * from budget.count_purchases()`)
	if err != nil {
		return IncorrectID, err
	}

	return total, err
}

//CreateNewPurchase creates new purchase and returns its id
func (r *PurchasesRepository) CreateNewPurchase(purchaseData *PurchaseWithGoods, sessionUUID uuid.UUID) (int, error) {
	var addedPurchaseID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&addedPurchaseID,
		`SELECT * from budget.create_new_purchase($1, $2, $3, $4, $5::uuid)`,
		purchaseData.TotalPrice,
		purchaseData.ShopID,
		//time.Unix(0, purchaseData.Date.),
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
		//time.Unix(0, purchaseData.Date),
		purchaseData.Comment,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return updatedPurchaseID, err
}

//IsPurchaseWithGoodsValid validates purchase
func (r *PurchasesRepository) IsPurchaseValid(data *Purchase) (bool, error) {
	if data.ShopID == nil {
		return false, errors.New("Validation failed. No shop provided")
	}

	if len(data.Comment) > 3000 {
		return false, errors.New("Validation failed. Purchase comment contains more than 3000 characters")
	}

	return true, nil
}

//IsPurchaseWithGoodsValid validates purchase with goods data
func (r *PurchasesRepository) IsPurchaseWithGoodsValid(data *PurchaseWithGoods) (bool, error) {
	if data.ShopID == nil {
		return false, errors.New("Validation failed. No shop provided")
	}

	if len(data.Comment) > 3000 {
		return false, errors.New("Validation failed. Purchase comment contains more than 3000 characters")
	}

	var totalPrice float32 = 0.0

	for _, item := range data.GoodsData {
		if len(item.Name) > 300 {
			return false, errors.New("Validation failed. Goods items name contains more than 300 characters")
		}

		if len(item.Comment) > 3000 {
			return false, errors.New("Validation failed. Goods items comment contains more than 3000 characters")
		}

		if item.Amount <= 0 {
			return false, errors.New("Validation failed. Goods items amount can not be <= 0")
		}

		if item.Price <= 0 {
			return false, errors.New("Validation failed. Goods items price can not be <= 0")
		}

		totalPrice += item.Price
	}

	if data.TotalPrice != totalPrice {
		return false, errors.New("Validation failed. Wrong total price")
	}

	return true, nil
}

package repos

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

//GoodsItem represents goods data item
type GoodsItem struct {
	Name       string `json:"name,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
	Comment    string `json:"comment,omitempty"`
	BarCode    string `json:"bar_code,omitempty"`

	Entity
}

//GoodsDetailsItem represents goods data details item
type GoodsDetailsItem struct {
	Amount      int     `json:"amount,omitempty"`
	Price       float32 `json:"price,omitempty"`
	PurchaseID  int     `json:"purchase_id,omitempty"`
	GoodsItemID int     `json:"goods_item_id,omitempty"`

	Entity
}

//GoodsItemWithDetails represents goods data item with details fields
type GoodsItemWithDetails struct {
	GoodsItem
	GoodsDetailsItem

	GoodsDetailsID int `json:"goods_details_id"`
	GoodsItemID    int `json:"goods_id"`
}

//SimpleGoodsItem represents goods data items for dropdown menus
type SimpleGoodsItem struct {
	SimpleEntity
}

//GoodsRepository is a repository for goods data
type GoodsRepository struct {
	EntityRepository
}

//GetEntityByID returns goods item by its id
func (r *GoodsRepository) GetEntityByID(id int) (GoodsItem, error) {
	var goodsItem GoodsItem

	err := pgxscan.Get(context.Background(), r.db, &goodsItem, `SELECT * from budget.get_goods_item_by_id($1)`, id)
	if err != nil {
		fmt.Println(err)
		return goodsItem, err
	}

	return goodsItem, err
}

//GetSlice returns slice of goods items
func (r *GoodsRepository) GetSlice(from int, count int) ([]*GoodsItem, error) {
	var goodsItems []*GoodsItem

	err := pgxscan.Select(context.Background(), r.db, &goodsItems, `SELECT * from budget.get_goods_items($1, $2)`, from, count)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return goodsItems, err
}

//GetTopGoodsItemsByName finds top goods items by name
func (r *GoodsRepository) GetTopGoodsItemsByName(name string) ([]*GoodsItem, error) {
	var goodsItems []*GoodsItem

	err := pgxscan.Select(context.Background(), r.db, &goodsItems, `SELECT * from budget.get_top_goods_items_by_name($1, $2)`, name, TopCounter)
	if err != nil {
		return nil, err
	}

	return goodsItems, err
}

//CreateNewGoodsItem creates new goods item and returns its id
func (r *GoodsRepository) CreateNewGoodsItem(goodsItemData *GoodsItem, sessionUUID uuid.UUID) (int, error) {
	var addedGoodsItemID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&addedGoodsItemID,
		`SELECT * from budget.create_new_goods_item($1, $2, $3, $4, $5::uuid)`,
		goodsItemData.Name,
		goodsItemData.CategoryID,
		goodsItemData.Comment,
		goodsItemData.BarCode,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return addedGoodsItemID, err
}

//RemoveEntityByID marks single goods item as removed in db
func (r *GoodsRepository) RemoveEntityByID(id int, uuid uuid.UUID) (int, error) {
	var removedGoodsItemID int

	err := pgxscan.Get(context.Background(), r.db, &removedGoodsItemID, `SELECT * from budget.remove_goods_item($1, $2::uuid)`, id, uuid)
	if err != nil {
		return IncorrectID, err
	}

	return removedGoodsItemID, err
}

//UpdateGoodsItem updates existing goods item and returns its id
func (r *GoodsRepository) UpdateGoodsItem(goodsItemData *GoodsItem, sessionUUID uuid.UUID) (int, error) {
	var updatedGoodsItemID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&updatedGoodsItemID,
		`SELECT * from budget.update_single_goods_item($1, $2, $3, $4, $5, $6::uuid)`,
		goodsItemData.ID,
		goodsItemData.Name,
		goodsItemData.CategoryID,
		goodsItemData.Comment,
		goodsItemData.BarCode,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return updatedGoodsItemID, err
}

//CreateNewGoodsDetailsItem creates new goods details item
func (r *GoodsRepository) CreateNewGoodsDetailsItem(goodsDetailsItemData *GoodsDetailsItem, sessionUUID uuid.UUID) (int, error) {
	var addedGoodsDetailsItemID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&addedGoodsDetailsItemID,
		`SELECT * from budget.create_new_goods_details($1, $2, $3, $4, $5::uuid)`,
		goodsDetailsItemData.Amount,
		goodsDetailsItemData.Price,
		goodsDetailsItemData.PurchaseID,
		goodsDetailsItemData.GoodsItemID,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return addedGoodsDetailsItemID, err
}

//RemoveGoodsDetails removes connection between purchase and goods item
func (r *GoodsRepository) RemoveGoodsDetails(goodsDetailsID, purchaseID, goodsItemID int, uuid uuid.UUID) (int, error) {
	var removedGoodsDetailsItemID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&removedGoodsDetailsItemID,
		`SELECT * from budget.remove_goods_details($1, $2, $3, $4::uuid)`,
		goodsDetailsID,
		purchaseID,
		goodsItemID,
		uuid)
	if err != nil {
		return IncorrectID, err
	}

	return removedGoodsDetailsItemID, err
}

//UpdateGoodsDetailsItem updates existing goods details item and returns its id
func (r *GoodsRepository) UpdateGoodsDetailsItem(goodsDetailsItemData *GoodsDetailsItem, sessionUUID uuid.UUID) (int, error) {
	var updatedGoodsDetailsItemID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&updatedGoodsDetailsItemID,
		`SELECT * from budget.update_goods_details($1, $2, $3, $4, $5, $6::uuid)`,
		goodsDetailsItemData.ID,
		goodsDetailsItemData.Amount,
		goodsDetailsItemData.Price,
		goodsDetailsItemData.PurchaseID,
		goodsDetailsItemData.GoodsItemID,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return updatedGoodsDetailsItemID, err
}

//GetGoodsDataForPurchase returns full goods data for purchase by its id
func (r *GoodsRepository) GetGoodsDataForPurchase(purchaseID int) ([]*GoodsItemWithDetails, error) {
	var goodsData []*GoodsItemWithDetails

	err := pgxscan.Select(context.Background(), r.db, &goodsData, `SELECT * from budget.get_goods_data_for_purchase($1)`, purchaseID)
	if err != nil {
		return nil, err
	}

	return goodsData, err
}

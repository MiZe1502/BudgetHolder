package repos

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
)

//GoodsItem represents goods data item
type GoodsItem struct {
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Comment    string `json:"comment"`
	BarCode    string `json:"bar_code"`

	Entity
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
func (r *ShopsRepository) GetTopGoodsItemsByName(name string) ([]*GoodsItem, error) {
	var goodsItems []*GoodsItem

	err := pgxscan.Select(context.Background(), r.db, &goodsItems, `SELECT * from budget.get_top_goods_items_by_name($1, $2)`, name, TopCounter)
	if err != nil {
		return nil, err
	}

	return goodsItems, err
}
package repos

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
)

//Category represents category data item
type Category struct {
	Name     string `json:"name"`
	Comment  string `json:"comment,omitempty"`
	ParentID *int   `json:"parent_id"`

	Categories []*Category `json:"categories,omitempty"`

	Entity
}

//SimpleCategory represents category data items for dropdown menus
type SimpleCategory struct {
	SimpleEntity
}

//CategoriesRepository is a repository for categories data
type CategoriesRepository struct {
	EntityRepository
}

//GetAllCategoriesAsChains returns all categories as tree structure
func (r *CategoriesRepository) GetAllCategoriesAsChains() ([]*Category, error) {
	var categories []*Category

	err := pgxscan.Select(context.Background(), r.db, &categories, `SELECT * from budget.get_all_goods_categories_chains()`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	catTree := r.CategoriesListToTree(categories)

	return catTree, err
}

//GetCategoryChainByParentID returns single categories chain
func (r *CategoriesRepository) GetCategoryChainByParentID(parentID int) ([]*Category, error) {
	var categories []*Category

	err := pgxscan.Select(context.Background(), r.db, &categories, `SELECT * from budget.get_goods_categories_chain_by_id($1)`, parentID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	catTree := r.CategoriesListToTree(categories)

	return catTree, err
}

//GetSingleCategoryByID returns single category by its id
func (r *CategoriesRepository) GetSingleCategoryByID(categoryID int) (*Category, error) {
	var category *Category

	err := pgxscan.Select(context.Background(), r.db, &category, `SELECT * from budget.get_goods_category_by_id($1)`, categoryID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return category, err
}

//GetSimpleCategoriesList returns list of simple categories for menus
func (r *CategoriesRepository) GetSimpleCategoriesList() ([]*SimpleCategory, error) {
	var categories []*SimpleCategory

	err := pgxscan.Select(context.Background(), r.db, &categories, `SELECT * from budget.get_simple_goods_categories()`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return categories, err
}

//CategoriesListToTree converts list categories structure into tree
func (r *CategoriesRepository) CategoriesListToTree(list []*Category) []*Category {
	m := map[int]int{}
	roots := []*Category{}

	for i := 0; i < len(list); i++ {
		m[list[i].ID] = i
		list[i].Categories = []*Category{}
	}

	for i := 0; i < len(list); i++ {
		node := list[i]
		if node.ParentID != nil {
			parentID := *node.ParentID
			list[m[parentID]].Categories = append(list[m[parentID]].Categories, node)
		} else {
			roots = append(roots, node)
		}
	}

	return roots
}
package repos

import (
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
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

//GetAllCategoriesAsTree returns all categories as tree structure
func (r *CategoriesRepository) GetAllCategoriesAsTree() ([]*Category, error) {
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
func (r *CategoriesRepository) GetCategoryChainByParentID(parentID *int) ([]*Category, error) {
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

//RemoveEntityByID marks single category as removed in db
func (r *CategoriesRepository) RemoveEntityByID(id int, uuid uuid.UUID) (int, error) {
	var removedCategoryID int

	err := pgxscan.Get(context.Background(), r.db, &removedCategoryID, `SELECT * from budget.remove_goods_category($1, $2::uuid)`, id, uuid)
	if err != nil {
		return IncorrectID, err
	}

	return removedCategoryID, err
}

//CreateNewCategory creates new goods category and returns its id
func (r *CategoriesRepository) CreateNewCategory(categoryData *Category, sessionUUID uuid.UUID) (int, error) {
	var addedCategoryID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&addedCategoryID,
		`SELECT * from budget.create_new_goods_category($1, $2, $3, $4::uuid)`,
		categoryData.Name,
		categoryData.Comment,
		categoryData.ParentID,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return addedCategoryID, err
}

//UpdateCategory updates existing goods category and returns its id
func (r *CategoriesRepository) UpdateCategory(categoryData *Category, sessionUUID uuid.UUID) (int, error) {
	var updatedCategoryID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&updatedCategoryID,
		`SELECT * from budget.update_goods_category($1, $2, $3, $4, $5::uuid)`,
		categoryData.ID,
		categoryData.Name,
		categoryData.Comment,
		categoryData.ParentID,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return updatedCategoryID, err
}

//IsCategoryValid validates category data
func (r *CategoriesRepository) IsCategoryValid(category *Category) (bool, error) {
	if len(category.Name) == 0 {
		return false, errors.New("Validation failed. No goods category name provided")
	}

	if len(category.Name) > 100 {
		return false, errors.New("Validation failed. Category name contains more than 100 characters")
	}

	if len(category.Comment) > 3000 {
		return false, errors.New("Validation failed. Category comment contains more than 100 characters")
	}

	return true, nil
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

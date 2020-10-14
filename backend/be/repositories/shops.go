package repos

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/pgxpool"
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
	Repository
}

//GetSlice returns slice of shops data
func (r *ShopsRepository) GetSlice(db *pgxpool.Pool, from int, to int) ([]*Shop, error) {
	var shops []*Shop

	err := pgxscan.Select(context.Background(), db, &shops, `SELECT id, name, comment from budget.get_shops($1, $2)`, from, to)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return shops, err
}

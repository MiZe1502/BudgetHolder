package repos

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

// User represents basic user data
type User struct {
	Login    string
	Password string

	Entity
}

// UserRepository provides methods to process user data
type UserRepository struct {
	EntityRepository
}

// GetUserByLogin returns single user with its encoded password
func (r *UserRepository) GetUserByLogin(login string) (User, error) {
	var user User

	err := pgxscan.Get(context.Background(), r.db, &user, `SELECT * from budget.get_user_by_login($1)`, login)
	if err != nil {
		return user, err
	}

	return user, err
}

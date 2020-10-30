package repos

import (
	"context"
	"errors"

	utils "../utils"

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
func (r *UserRepository) getUserByLogin(login string) (User, error) {
	var user User

	err := pgxscan.Get(context.Background(), r.db, &user, `SELECT * from budget.get_user_by_login($1)`, login)
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *UserRepository) createUserSession(login, ip string) (string, error) {
	uuid := utils.GetNewUUID()

	err := pgxscan.Get(context.Background(), r.db, &uuid, `SELECT * from budget.create_session($1, $2, $3)`, login, ip, uuid)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

// ProcessUserAuth contains processing login logic:
// get user data from db by login
// check password and its hash
// create session and save it to db
// create websocket connection and add it to Hub
// create jwt token and return it as response
func (r *UserRepository) ProcessUserAuth(login string, password string, ip string) (string, error) {
	user, err := r.getUserByLogin(login)
	if err != nil {
		r.log.Error(err.Error())
		return "", err
	}

	isPasswordCorrect := utils.ComparePasswords(user.Password, password)
	if !isPasswordCorrect {
		err := errors.New("Incorrect password: " + password + " for user: " + login)
		r.log.Error(err.Error())
		return "", err
	}

	uuid, err := r.createUserSession(login, ip)
	if err != nil {
		r.log.Error(err.Error())
		return "", err
	}

	r.log.Info("created session: " + uuid + " for user: " + login)

	parsedUUID, err := utils.GetUUIDfromString(uuid)
	if err != nil {
		r.log.Error(err.Error())
		return "", err
	}

	token, err := r.token.CreateNewToken(parsedUUID)
	if err != nil {
		r.log.Error(err.Error())
		return "", err
	}

	return token, nil
}

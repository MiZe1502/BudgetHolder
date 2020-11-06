package repos

import (
	"context"
	"errors"

	utils "../utils"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

// UserGroup stores user group data
type UserGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Entity
}

// User represents basic user data
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`

	Entity
}

// FullUser contains all data about user profile, login, pass etc.
type FullUser struct {
	User

	Name        string `json:"name"`
	Surname     string `json:"surname"`
	ImagePath   string `json:"path"`
	Description string `json:"description"`
	GroupName   string `json:"groupname"`
}

// UserContext stores user data for request context and web socket connections
type UserContext struct {
	SessionUUID uuid.UUID
	UserID      int
	UserGroupID int
	IP          string
}

// UserRepository provides methods to process user data
type UserRepository struct {
	EntityRepository
}

// IsUserValid validates user data
func (r *UserRepository) IsUserValid(user *FullUser) (bool, error) {
	if len(user.Login) == 0 {
		return false, errors.New("Validation failed. No user login provided")
	}

	if len(user.Login) > 100 {
		return false, errors.New("Validation failed. User login contains more than 100 characters")
	}

	if len(user.Password) < 8 {
		return false, errors.New("Validation failed. User password is too short")
	}

	if len(user.Name) > 100 {
		return false, errors.New("Validation failed. Too long user name")
	}

	if len(user.Name) > 100 {
		return false, errors.New("Validation failed. Too long user name")
	}

	if len(user.Surname) > 100 {
		return false, errors.New("Validation failed. Too long user surname")
	}

	if len(user.Description) > 3000 {
		return false, errors.New("Validation failed. Too long user description")
	}

	return true, nil
}

// GetUserGroupIDByUserID returns user group id by user id from this group
func (r *UserRepository) GetUserGroupIDByUserID(userID int) (int, error) {
	var groupID int

	err := pgxscan.Get(context.Background(), r.db, &groupID, `SELECT * from budget.get_user_group_id_by_user_id($1)`, userID)
	if err != nil {
		return IncorrectID, err
	}

	return userID, err
}

// GetUserIDBySessionUUID returns userId by its current session UUID
func (r *UserRepository) GetUserIDBySessionUUID(sessionUUID uuid.UUID) (int, error) {
	var userID int

	err := pgxscan.Get(context.Background(), r.db, &userID, `SELECT * from budget.get_user_id_by_session_uuid($1::uuid)`, sessionUUID)
	if err != nil {
		return IncorrectID, err
	}

	return userID, err
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
	uuid, err := utils.GetUUIDfromString(utils.GetNewUUID())
	if err != nil {
		return "", err
	}

	err = pgxscan.Get(context.Background(), r.db, &uuid, `SELECT * from budget.create_session($1, $2, $3::uuid)`, login, ip, uuid)
	if err != nil {
		return "", err
	}

	return utils.DecodeUUIDIntoString(uuid), nil
}

// getUserGroupByName returns single user with its encoded password
func (r *UserRepository) getUserGroupByName(groupName string) (UserGroup, error) {
	var group UserGroup

	err := pgxscan.Get(context.Background(), r.db, &group, `SELECT * from budget.get_user_group_by_name($1)`, groupName)
	if err != nil {
		return group, err
	}

	return group, err
}

//CreateNewUser creates new user and returns its id
func (r *UserRepository) CreateNewUser(userData *FullUser) (int, error) {
	var newUserID int

	userGroup, err := r.getUserGroupByName(userData.GroupName)
	if err != nil {
		return IncorrectID, err
	}

	pwd, err := utils.HashPasswordWithSalt(userData.Password)
	if err != nil {
		return IncorrectID, err
	}

	err = pgxscan.Get(context.Background(),
		r.db,
		&newUserID,
		`SELECT * from budget.create_user($1, $2, $3, $4, $5, $6, $7)`,
		userData.Login,
		pwd,
		userData.Name,
		userData.Surname,
		userData.ImagePath,
		userData.Description,
		userGroup.ID)
	if err != nil {
		return IncorrectID, err
	}

	return newUserID, nil
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

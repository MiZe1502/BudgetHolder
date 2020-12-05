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
	Description string `json:"description,omitempty"`

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

	Name             string `json:"name,omitempty"`
	Surname          string `json:"surname,omitempty"`
	PathToPhoto      string `json:"path_to_photo,omitempty"`
	Description      string `json:"description,omitempty"`
	GroupName        string `json:"group_name"`
	GroupDescription string `json:"group_description"`
}

// UpdatedUser is extended struct to store data for updated user
type UpdatedUser struct {
	FullUser

	IsPasswordChanged bool `json:"is_password_changed,omitempty"`
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

// IsUserGroupValid validates user group data
func (r *UserRepository) IsUserGroupValid(group *UserGroup) (bool, error) {
	if len(group.Name) == 0 {
		return false, errors.New("Validation failed. No user group name provided")
	}

	if len(group.Name) > 100 {
		return false, errors.New("Validation failed. User group name contains more than 100 characters")
	}

	if len(group.Description) > 3000 {
		return false, errors.New("Validation failed. Too long user group description")
	}

	return true, nil
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

// GetUserLoginByID returns user login by its id
func (r *UserRepository) GetUserLoginByID(userID int) (string, error) {
	var userLogin string

	err := pgxscan.Get(context.Background(), r.db, &userLogin, `SELECT * from budget.get_user_by_id($1)`, userID)
	if err != nil {
		return userLogin, err
	}

	return userLogin, err
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

//CloseUserSession closes user session
func (r *UserRepository) CloseUserSession(sessionUUID uuid.UUID) (uuid.UUID, error) {
	var closedSessionUUID uuid.UUID

	err := pgxscan.Get(context.Background(), r.db, &closedSessionUUID, `SELECT * from budget.close_session($1::uuid)`, sessionUUID)
	if err != nil {
		return closedSessionUUID, err
	}

	return closedSessionUUID, err
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

//CreateNewUserGroup creates new user and returns its id
func (r *UserRepository) CreateNewUserGroup(groupData *UserGroup, sessionUUID uuid.UUID) (int, error) {
	var newID int

	err := pgxscan.Get(context.Background(),
		r.db, &newID,
		`SELECT * from budget.create_user_group($1, $2, $3::uuid)`,
		groupData.Name,
		groupData.Description,
		sessionUUID.String())
	if err != nil {
		return IncorrectID, err
	}

	return newID, err
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
		userData.PathToPhoto,
		userData.Description,
		userGroup.ID)
	if err != nil {
		return IncorrectID, err
	}

	return newUserID, nil
}

//GetFullUserInfo returns full user data
func (r *UserRepository) GetFullUserInfo(login string) (FullUser, error) {
	var user FullUser

	err := pgxscan.Get(context.Background(), r.db, &user, `SELECT * from budget.get_full_user_info($1)`, login)
	if err != nil {
		return user, err
	}

	return user, err
}

//RemoveUserByLogin marks user as removed by login
func (r *UserRepository) RemoveUserByLogin(login string) (int, error) {
	var userID int

	err := pgxscan.Get(context.Background(), r.db, &userID, `SELECT * from budget.remove_user($1)`, login)
	if err != nil {
		return IncorrectID, err
	}

	return userID, err
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

// ActualizeUserLastOnlineByLogin updates last online time for user by its login
func (r *UserRepository) ActualizeUserLastOnlineByLogin(userLogin string) (int, error) {
	var userID int

	err := pgxscan.Get(context.Background(), r.db, &userID, `SELECT * from budget.actualize_user_online($1)`, userLogin)
	if err != nil {
		return IncorrectID, err
	}

	return userID, err
}

//UpdateUser updates existing user and returns its id
func (r *UserRepository) UpdateUser(oldLogin string, userData *UpdatedUser) (int, error) {
	var updatedUserID int

	err := pgxscan.Get(context.Background(),
		r.db,
		&updatedUserID,
		`SELECT * from budget.update_user($1, $2, $3, $4, $5, $6, $7, $8)`,
		oldLogin,
		userData.Login,
		userData.IsPasswordChanged,
		userData.Password,
		userData.PathToPhoto,
		userData.Name,
		userData.Surname,
		userData.Description)
	if err != nil {
		return IncorrectID, err
	}

	return updatedUserID, err
}

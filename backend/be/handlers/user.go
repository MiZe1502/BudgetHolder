package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	env "../env"
	repos "../repositories"
	utils "../utils"
)

func createNewUserGroupHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createNewUserGroupHandler: start")

		env.Logger.Info("createNewUserGroupHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createNewUserGroupHandler: getting user group data from request")

		groupData := &repos.UserGroup{}
		err := json.NewDecoder(r.Body).Decode(groupData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createNewUserGroupHandler: init user repository")

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createNewUserGroupHandler: validate user group")

		isValid, err := repo.IsUserGroupValid(groupData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusBadRequest)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		if !isValid {
			msg := utils.MessageError(utils.Message(false, "Smth went wrong. Validation failed without error"), http.StatusBadRequest)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createNewUserGroupHandler: adding user group with name: " + groupData.Name)

		user := r.Context().Value("userCtx").(*repos.UserContext)
		newGroupID, err := repo.CreateNewUserGroup(groupData, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createNewUserGroupHandler: added user group with name: " + groupData.Name + " | id: " + fmt.Sprint(newGroupID))

		groupData.ID = newGroupID
		groupDataJSON, err := json.Marshal(groupData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), groupDataJSON), env.Logger)
	}
}

func createNewUserHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createNewUserHandler")

		env.Logger.Info("check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("getting user data from request")

		userData := &repos.FullUser{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("init user repository")

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("validate user: login: " + userData.Login)

		isValid, err := repo.IsUserValid(userData, true)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusBadRequest)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		if !isValid {
			msg := utils.MessageError(utils.Message(false, "Smth went wrong. Validation failed without error"), http.StatusBadRequest)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("adding user: login: " + userData.Login)

		newUserID, err := repo.CreateNewUser(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("added user: login: " + userData.Login + " | id: " + fmt.Sprint(newUserID))

		userDataJSON, err := json.Marshal(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), userDataJSON), env.Logger)
	}
}

func createAuthHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createAuthHandler")

		env.Logger.Info("check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("getting user data from request")

		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("processing user auth: login: " + userData.Login + " | pass: " + userData.Password)

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		token, err := repo.ProcessUserAuth(userData.Login, userData.Password, r.RemoteAddr)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("got token: " + token + " for user: " + userData.Login)

		utils.Respond(w, utils.Message(true, token), env.Logger)
	}
}

func createActualizeUserLastOnlineHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createActualizeUserLastOnlineHandler")

		env.Logger.Info("createActualizeUserLastOnlineHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createActualizeUserLastOnlineHandler: getting user data from request")

		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createActualizeUserLastOnlineHandler: processing user with login: " + userData.Login)

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		userID, err := repo.ActualizeUserLastOnlineByLogin(userData.Login)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createActualizeUserLastOnlineHandler: updated last online datetime for user with id: " + fmt.Sprint(userID))

		utils.Respond(w, utils.Message(true, fmt.Sprint(userID)), env.Logger)
	}
}

func createGetFullUserInfoHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetFullUserInfoHandler")

		env.Logger.Info("createGetFullUserInfoHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetFullUserInfoHandler: getting user data from request")

		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetFullUserInfoHandler: processing user with login: " + userData.Login)

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		user, err := repo.GetFullUserInfo(userData.Login)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetFullUserInfoHandler: marshalling user data: login: " + userData.Login + " | id: " + fmt.Sprint(userData.ID))

		userJSON, err := json.Marshal(user)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), userJSON), env.Logger)
	}
}

func createRemoveUserHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createRemoveUserHandler: start")

		env.Logger.Info("createRemoveUserHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveUserHandler: getting data from request")

		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveUserHandler: init user repository")

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createRemoveUserHandler: removing user with login: " + userData.Login)

		removedUserID, err := repo.RemoveUserByLogin(userData.Login)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveUserHandler: successfully removed user with id: " + fmt.Sprint(removedUserID))

		utils.Respond(w, utils.Message(true, fmt.Sprint(removedUserID)), env.Logger)
	}
}

func createCloseUserSessionHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createCloseUserSessionHandler: start")

		env.Logger.Info("createCloseUserSessionHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createCloseUserSessionHandler: init user repository")

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		user := r.Context().Value("userCtx").(*repos.UserContext)

		env.Logger.Info("createCloseUserSessionHandler: closing user with UUID: " + user.SessionUUID.String())

		closedSessionUUID, err := repo.CloseUserSession(user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveUserHandler: successfully closed user session with UUID: " + fmt.Sprint(closedSessionUUID))

		utils.Respond(w, utils.Message(true, "OK"), env.Logger)
	}
}

func createUpdateUserHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createUpdateUserHandler: start")

		env.Logger.Info("createUpdateUserHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateUserHandler: getting user data from request")

		userData := &repos.UpdatedUser{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateUserHandler: init user repository")

		var repo repos.UserRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		user := r.Context().Value("userCtx").(*repos.UserContext)

		env.Logger.Info("getting user login by id: " + fmt.Sprint(user.UserID))

		oldUserLogin, err := repo.GetUserLoginByID(user.UserID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		//TODO: Implement validation

		env.Logger.Info("createUpdateUserHandler: validate user: old login: " + oldUserLogin + " | new login: " + userData.Login)

		var userDataForValidation *repos.FullUser

		jsonedData, err := json.Marshal(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		err = json.Unmarshal(jsonedData, &userDataForValidation)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		isValid, err := repo.IsUserValid(userDataForValidation, userData.IsPasswordChanged)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusBadRequest)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		if !isValid {
			msg := utils.MessageError(utils.Message(false, "Smth went wrong. Validation failed without error"), http.StatusBadRequest)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("updating user: old login: " + oldUserLogin + " | new login: " + userData.Login)

		updatedUserID, err := repo.UpdateUser(oldUserLogin, userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("user updated: login: " + userData.Login + " | id: " + fmt.Sprint(updatedUserID))

		utils.Respond(w, utils.Message(true, ""), env.Logger)
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	env "../env"
	repos "../repositories"
	utils "../utils"
)

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

		isValid, err := repo.IsUserValid(userData)

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
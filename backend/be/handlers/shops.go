package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	env "../env"
	repos "../repositories"
	utils "../utils"
)

func createGetShopsSliceHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetShopsSliceHandler: start")

		env.Logger.Info("createGetShopsSliceHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopsSliceHandler: getting data from request")

		sliceData := &utils.SliceData{}

		err := json.NewDecoder(r.Body).Decode(sliceData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopsSliceHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetShopsSliceHandler: getting slice of shops, from:" + fmt.Sprint(sliceData.From) + ", count: " + fmt.Sprint(sliceData.Count))

		shops, err := repo.GetSlice(sliceData.From, sliceData.Count)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopsSliceHandler: marshalling slice of shops")

		shopsJSON, err := json.Marshal(shops)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), shopsJSON), env.Logger)
	}
}

func createGetTopShopsByNameHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetTopShopsByNameHandler: start")

		env.Logger.Info("createGetTopShopsByNameHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetTopShopsByNameHandler: getting data from request")

		shopData := &repos.SimpleShop{}

		err := json.NewDecoder(r.Body).Decode(shopData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetTopShopsByNameHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetTopShopsByNameHandler: getting list of top shops for name: " + shopData.Name)

		shops, err := repo.GetTopShopsByName(shopData.Name)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetTopShopsByNameHandler: marshalling list of top shops")

		shopsJSON, err := json.Marshal(shops)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), shopsJSON), env.Logger)
	}
}

func createAddNewShopHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createAddNewShopHandler: start")

		env.Logger.Info("createAddNewShopHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createAddNewShopHandler: getting data from request")

		shopData := &repos.Shop{}

		err := json.NewDecoder(r.Body).Decode(shopData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createAddNewShopHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createAddNewShopHandler: creating new shop")

		user := r.Context().Value("userCtx").(*repos.UserContext)
		addedShopID, err := repo.CreateNewShop(shopData, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createAddNewShopHandler: created new shop with id: " + fmt.Sprint(addedShopID))

		shopData.ID = addedShopID
		shopsJSON, err := json.Marshal(shopData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), shopsJSON), env.Logger)
	}
}

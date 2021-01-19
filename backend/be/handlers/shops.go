package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

		from, err := strconv.Atoi(r.URL.Query().Get("from"))
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "createGetShopsSliceHandler: Invalid parsing data from request"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "createGetShopsSliceHandler: Invalid parsing data from request"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopsSliceHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetShopsSliceHandler: getting slice of shops, from:" + fmt.Sprint(from) + ", count: " + fmt.Sprint(count))

		shops, err := repo.GetSlice(from, count)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopsSliceHandler: getting total of shops")

		total, err := repo.CountTotal()
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopsSliceHandler: marshalling slice of shops")

		data := &utils.TableData{}
		data.Total = total
		data.Data = shops

		shopsJSON, err := json.Marshal(data)
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

		env.Logger.Info("createAddNewShopHandler: validate new shop")

		isValid, err := repo.IsShopValid(shopData)

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

func createUpdateShopHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createUpdateShopHandler: start")

		env.Logger.Info("createUpdateShopHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateShopHandler: getting data from request")

		shopData := &repos.Shop{}

		err := json.NewDecoder(r.Body).Decode(shopData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateShopHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createUpdateShopHandler: validate shop with id: " + fmt.Sprint(shopData.ID))

		isValid, err := repo.IsShopValid(shopData)

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

		env.Logger.Info("createUpdateShopHandler: updating shop with id: " + fmt.Sprint(shopData.ID))

		user := r.Context().Value("userCtx").(*repos.UserContext)
		updatedShopID, err := repo.UpdateShop(shopData, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateShopHandler: updated shop with id: " + fmt.Sprint(updatedShopID))

		shopsJSON, err := json.Marshal(shopData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), shopsJSON), env.Logger)
	}
}

func createRemoveShopHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createRemoveShopHandler: start")

		env.Logger.Info("createRemoveShopHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveShopHandler: getting data from request")

		shopData := &repos.SimpleShop{}

		err := json.NewDecoder(r.Body).Decode(shopData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveShopHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createRemoveShopHandler: removing shop with id: " + fmt.Sprint(shopData.ID))

		user := r.Context().Value("userCtx").(*repos.UserContext)
		removedShopID, err := repo.RemoveEntityByID(shopData.ID, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveShopHandler: successfully removed shop with id: " + fmt.Sprint(removedShopID))

		utils.Respond(w, utils.Message(true, fmt.Sprint(removedShopID)), env.Logger)
	}
}

func createGetShopByIDHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetShopByIDHandler: start")

		env.Logger.Info("createGetShopByIDHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopByIDHandler: getting data from request")

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid parsing data from request"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetShopByIDHandler: init shops repository")

		var repo repos.ShopsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createRemoveShopHandler: getting shop with id: " + fmt.Sprint(id))

		shop, err := repo.GetEntityByID(id)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		shopJSON, err := json.Marshal(shop)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), shopJSON), env.Logger)
	}
}

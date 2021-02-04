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

func createGetGoodsSliceHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetGoodsSliceHandler: start")

		env.Logger.Info("createGetGoodsSliceHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsSliceHandler: getting data from request")

		from, err := strconv.Atoi(r.URL.Query().Get("from"))
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "createGetGoodsSliceHandler: Invalid parsing data from request"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "createGetGoodsSliceHandler: Invalid parsing data from request"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsSliceHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		user := r.Context().Value("userCtx").(*repos.UserContext)

		goods, err := repo.GetSlice(from, count, user.UserID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsSliceHandler: init categories repository")

		var catRepo repos.CategoriesRepository

		catRepo.SetDb(env.Db)
		catRepo.SetLogger(env.Logger)
		catRepo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetGoodsSliceHandler: getting categories for goods items")

		categories, err := catRepo.GetSimpleCategoriesList()
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		for _, g := range goods {
			for _, c := range categories {
				if g.CategoryID == c.ID {
					g.Category = &repos.SimpleEntity{ID: g.CategoryID, Name: c.Name}
				}
			}
		}

		env.Logger.Info("createGetGoodsSliceHandler: getting total of goods items")

		total, err := repo.CountTotal()
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsSliceHandler: marshalling slice of goods items")

		data := &utils.TableData{}
		data.Total = total
		data.Data = goods

		goodsJSON, err := json.Marshal(data)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), goodsJSON), env.Logger)
	}
}

func createGetGoodsItemByIDHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetGoodsItemByIDHandler: start")

		env.Logger.Info("createGetGoodsItemByIDHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsItemByIDHandler: getting data from request")

		goodsItemData := &repos.Entity{}

		err := json.NewDecoder(r.Body).Decode(goodsItemData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsItemByIDHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetGoodsItemByIDHandler: getting goods item with id: " + fmt.Sprint(goodsItemData.ID))

		goodsItem, err := repo.GetEntityByID(goodsItemData.ID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		goodsItemJSON, err := json.Marshal(goodsItem)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), goodsItemJSON), env.Logger)
	}
}

func createGetSimpleGoodsItemsListHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetSimpleGoodsItemsListHandler: start")

		env.Logger.Info("createGetSimpleGoodsItemsListHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSimpleGoodsItemsListHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetSimpleGoodsItemsListHandler: getting list of simple goods items")

		goodsItems, err := repo.GetSimpleGoodsItemsList()
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSimpleGoodsItemsListHandler: marshalling list of simple goods items")

		goodsItemsJSON, err := json.Marshal(goodsItems)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), goodsItemsJSON), env.Logger)
	}
}

func createRemoveGoodsItemHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createRemoveGoodsItemHandler: start")

		env.Logger.Info("createRemoveGoodsItemHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveGoodsItemHandler: getting data from request")

		goodsData := &repos.Entity{}

		err := json.NewDecoder(r.Body).Decode(goodsData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveGoodsItemHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createRemoveGoodsItemHandler: removing goods item with id: " + fmt.Sprint(goodsData.ID))

		user := r.Context().Value("userCtx").(*repos.UserContext)
		removedGoodsItemID, err := repo.RemoveEntityByID(goodsData.ID, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveGoodsItemHandler: successfully removed goods item with id: " + fmt.Sprint(removedGoodsItemID))

		utils.Respond(w, utils.Message(true, fmt.Sprint(removedGoodsItemID)), env.Logger)
	}
}

func createAddNewGoodsItemHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createAddNewGoodsItemHandler: start")

		env.Logger.Info("createAddNewGoodsItemHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createAddNewGoodsItemHandler: getting data from request")

		goodsItemData := &repos.GoodsItem{}

		err := json.NewDecoder(r.Body).Decode(goodsItemData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createAddNewGoodsItemHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createAddNewGoodsItemHandler: validate new goods item")

		isValid, err := repo.IsGoodsItemValid(goodsItemData)

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

		env.Logger.Info("createAddNewGoodsItemHandler: creating new goods item")

		user := r.Context().Value("userCtx").(*repos.UserContext)
		addedGoodsItemID, err := repo.CreateNewGoodsItem(goodsItemData, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createAddNewGoodsItemHandler: created new goods item with id: " + fmt.Sprint(addedGoodsItemID))

		goodsItemData.ID = addedGoodsItemID
		goodsItemJSON, err := json.Marshal(goodsItemData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), goodsItemJSON), env.Logger)
	}
}

func createUpdateGoodsItemHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createUpdateGoodsItemHandler: start")

		env.Logger.Info("createUpdateGoodsItemHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateGoodsItemHandler: getting data from request")

		goodsItemData := &repos.GoodsItem{}

		err := json.NewDecoder(r.Body).Decode(goodsItemData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateGoodsItemHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createUpdateGoodsItemHandler: validate goods item with id: " + fmt.Sprint(goodsItemData.ID))

		isValid, err := repo.IsGoodsItemValid(goodsItemData)

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

		env.Logger.Info("createUpdateGoodsItemHandler: updating goods item with id: " + fmt.Sprint(goodsItemData.ID))

		user := r.Context().Value("userCtx").(*repos.UserContext)
		updatedGoodsItemID, err := repo.UpdateGoodsItem(goodsItemData, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createUpdateGoodsItemHandler: updated goods item with id: " + fmt.Sprint(updatedGoodsItemID))

		goodsItemJSON, err := json.Marshal(goodsItemData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), goodsItemJSON), env.Logger)
	}
}

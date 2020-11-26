package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	env "../env"
	repos "../repositories"
	utils "../utils"
)

func createGetGoodsCategoriesTreeHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetGoodsCategoriesTreeHandler: start")

		env.Logger.Info("createGetGoodsCategoriesTreeHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsCategoriesTreeHandler: init categories repository")

		var repo repos.CategoriesRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetGoodsCategoriesTreeHandler: getting categories tree")

		categories, err := repo.GetAllCategoriesAsTree()
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetGoodsCategoriesTreeHandler: marshalling categories tree")

		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), categoriesJSON), env.Logger)
	}
}

func createGetCategoryChainByParentIDHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetCategoryChainByParentIDHandler: start")

		env.Logger.Info("createGetCategoryChainByParentIDHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetCategoryChainByParentIDHandler: getting data from request")

		categoryData := &repos.Category{}

		err := json.NewDecoder(r.Body).Decode(categoryData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetCategoryChainByParentIDHandler: init categories repository")

		var repo repos.CategoriesRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetCategoryChainByParentIDHandler: getting category chain by parentID: " + fmt.Sprint(categoryData.ParentID))

		categories, err := repo.GetCategoryChainByParentID(categoryData.ParentID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetCategoryChainByParentIDHandler: marshalling categories chain")

		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), categoriesJSON), env.Logger)
	}
}

func createGetSingleCategoryByIDHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetSingleCategoryByIDHandler: start")

		env.Logger.Info("createGetSingleCategoryByIDHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSingleCategoryByIDHandler: getting data from request")

		categoryData := &repos.SimpleCategory{}

		err := json.NewDecoder(r.Body).Decode(categoryData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSingleCategoryByIDHandler: init categories repository")

		var repo repos.CategoriesRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetSingleCategoryByIDHandler: getting category by ID: " + fmt.Sprint(categoryData.ID))

		category, err := repo.GetEntityByID(categoryData.ID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSingleCategoryByIDHandler: marshalling category")

		categoryJSON, err := json.Marshal(category)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), categoryJSON), env.Logger)

	}
}

func createGetSimpleCategoriesListHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetSimpleCategoriesListHandler: start")

		env.Logger.Info("createGetSimpleCategoriesListHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSimpleCategoriesListHandler: init categories repository")

		var repo repos.CategoriesRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetSimpleCategoriesListHandler: getting simple categories list")

		categories, err := repo.GetSimpleCategoriesList()
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetSimpleCategoriesListHandler: marshalling simple categories list")

		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), categoriesJSON), env.Logger)
	}
}

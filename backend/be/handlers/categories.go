package handlers

import (
	"encoding/json"
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

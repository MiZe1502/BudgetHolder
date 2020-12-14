package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	env "../env"
	repos "../repositories"
	utils "../utils"
)

func createGetPurchasesWithGoodsDataSliceHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: start")

		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: check request method: " + r.Method)

		if r.Method != "GET" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: getting data from request")

		sliceData := &utils.SliceData{}

		err := json.NewDecoder(r.Body).Decode(sliceData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: init purchases repository")

		var purRepo repos.PurchasesRepository

		purRepo.SetDb(env.Db)
		purRepo.SetLogger(env.Logger)
		purRepo.SetTokenGenerator(env.Token)

		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: getting slice of purchases, from:" + fmt.Sprint(sliceData.From) + ", count: " + fmt.Sprint(sliceData.Count))

		user := r.Context().Value("userCtx").(*repos.UserContext)

		purchases, err := purRepo.GetSlice(sliceData.From, sliceData.Count, user.UserID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: init goods repository")

		var goodsRepo repos.GoodsRepository

		goodsRepo.SetDb(env.Db)
		goodsRepo.SetLogger(env.Logger)
		goodsRepo.SetTokenGenerator(env.Token)

		//TODO: we can use single request to db to get all purchases by array of ids
		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: getting goods data for purchases")
		for _, p := range purchases {
			goodsData, err := goodsRepo.GetGoodsDataForPurchase(p.ID)
			if err != nil {
				msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
				utils.RespondError(w, msg, env.Logger)
				return
			}
			p.GoodsData = goodsData
		}

		env.Logger.Info("createGetPurchasesWithGoodsDataSliceHandler: marshalling slice of purchases with goods data")

		purchasesJSON, err := json.Marshal(purchases)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), purchasesJSON), env.Logger)
	}
}

func createRemovePurchasesWithGoodsDetailsHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createRemovePurchasesWithGoodsDetailsHandler: start")

		env.Logger.Info("createRemovePurchasesWithGoodsDetailsHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		purchaseData := &repos.Entity{}

		err := json.NewDecoder(r.Body).Decode(purchaseData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemovePurchasesWithGoodsDetailsHandler: init purchase repository")

		var repo repos.PurchasesRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createRemovePurchasesWithGoodsDetailsHandler: removing purchase with id: " + fmt.Sprint(purchaseData.ID))

		user := r.Context().Value("userCtx").(*repos.UserContext)
		removedPurchaseID, err := repo.RemoveEntityByID(purchaseData.ID, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemovePurchasesWithGoodsDetailsHandler: successfully removed purchase with id: " + fmt.Sprint(removedPurchaseID))

		utils.Respond(w, utils.Message(true, fmt.Sprint(removedPurchaseID)), env.Logger)
	}
}

func createRemoveGoodsDetailsItemHandler(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createRemoveGoodsDetailsItemHandler: start")

		env.Logger.Info("createRemoveGoodsDetailsItemHandler: check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		goodsDetailsData := &repos.Entity{}

		err := json.NewDecoder(r.Body).Decode(goodsDetailsData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveGoodsDetailsItemHandler: init goods repository")

		var repo repos.GoodsRepository

		repo.SetDb(env.Db)
		repo.SetLogger(env.Logger)
		repo.SetTokenGenerator(env.Token)

		env.Logger.Info("createRemoveGoodsDetailsItemHandler: removing goods details with id: " + fmt.Sprint(goodsDetailsData.ID))

		user := r.Context().Value("userCtx").(*repos.UserContext)
		removedGoodsDetailsID, err := repo.RemoveGoodsDetails(goodsDetailsData.ID, user.SessionUUID)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		env.Logger.Info("createRemoveGoodsDetailsItemHandler: successfully removed goods details with id: " + fmt.Sprint(removedGoodsDetailsID))

		utils.Respond(w, utils.Message(true, fmt.Sprint(removedGoodsDetailsID)), env.Logger)
	}
}

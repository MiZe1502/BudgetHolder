package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	repos "./repositories"
	utils "./utils"
	"github.com/justinas/alice"
)

func createNewUserHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.logger.Info("createNewUserHandler")

		env.logger.Info("check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("getting user data from request")

		userData := &repos.FullUser{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("init user repository")

		var repo repos.UserRepository

		repo.SetDb(env.db)
		repo.SetLogger(env.logger)
		repo.SetTokenGenerator(env.token)

		env.logger.Info("validate user: login: " + userData.Login)

		isValid, err := repo.IsUserValid(userData)

		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusBadRequest)
			utils.RespondError(w, msg, env.logger)
			return
		}

		if !isValid {
			msg := utils.MessageError(utils.Message(false, "Smth went wrong. Validation failed without error"), http.StatusBadRequest)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("adding user: login: " + userData.Login)

		newUserID, err := repo.CreateNewUser(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("added user: login: " + userData.Login + " | id: " + fmt.Sprint(newUserID))

		userDataJSON, err := json.Marshal(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), userDataJSON), env.logger)
	}
}

func createAuthHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.logger.Info("createAuthHandler")

		env.logger.Info("check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("getting user data from request")

		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("processing user auth: login: " + userData.Login + " | pass: " + userData.Password)

		var repo repos.UserRepository

		repo.SetDb(env.db)
		repo.SetLogger(env.logger)
		repo.SetTokenGenerator(env.token)

		token, err := repo.ProcessUserAuth(userData.Login, userData.Password, r.RemoteAddr)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("got token: " + token + " for user: " + userData.Login)

		utils.Respond(w, utils.Message(true, token), env.logger)
	}
}

func createTestMessageHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		user := r.Context().Value("userCtx").(*repos.UserContext)

		// Print the message to the console
		fmt.Printf("message: %s\n", fmt.Sprint(user.UserID))

		env.hub.send <- body
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

func initHandlers(env *Env) {
	middlewareChain := alice.New(createMiddleware(env, Trace),
		createMiddleware(env, Logger),
		createMiddleware(env, CheckSession))

	http.Handle("/", middlewareChain.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/ws", middlewareChain.Then(http.HandlerFunc(createWsHandler(env))))

	http.Handle("/api/v1/user/auth", middlewareChain.Then(http.HandlerFunc(createAuthHandler(env))))
	http.Handle("/api/v1/user/new", middlewareChain.Then(http.HandlerFunc(createNewUserHandler(env))))

	http.Handle("/message", middlewareChain.Then(http.HandlerFunc(createTestMessageHandler(env))))

	http.ListenAndServe(":8080", nil)

}

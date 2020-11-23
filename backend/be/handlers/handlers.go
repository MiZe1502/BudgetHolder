package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	env "../env"
	wshub "../hub"
	repos "../repositories"
	"github.com/justinas/alice"
)

func createTestMessageHandler(env *env.Env, hub *wshub.Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		user := r.Context().Value("userCtx").(*repos.UserContext)

		// Print the message to the console
		fmt.Printf("message: %s\n", fmt.Sprint(user.UserID))

		hub.Send <- body
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

// InitHandlers initialize all endpoints with middlewares and handlers
func InitHandlers(env *env.Env, hub *wshub.Hub) {
	middlewareChain := alice.New(createMiddleware(env, Trace),
		createMiddleware(env, Logger),
		createMiddleware(env, CheckSession))

	http.Handle("/", middlewareChain.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/ws", middlewareChain.Then(http.HandlerFunc(createWsHandler(env, hub))))

	http.Handle("/api/v1/user/auth", middlewareChain.Then(http.HandlerFunc(createAuthHandler(env))))
	http.Handle("/api/v1/user/new", middlewareChain.Then(http.HandlerFunc(createNewUserHandler(env))))

	http.Handle("/api/v1/shops/slice", middlewareChain.Then(http.HandlerFunc(createGetShopsSliceHandler(env))))
	http.Handle("/api/v1/shops/top", middlewareChain.Then(http.HandlerFunc(createGetTopShopsByNameHandler(env))))
	http.Handle("/api/v1/shops/new", middlewareChain.Then(http.HandlerFunc(createAddNewShopHandler(env))))
	http.Handle("/api/v1/shops/remove", middlewareChain.Then(http.HandlerFunc(createRemoveShopHandler(env))))
	http.Handle("/api/v1/shops/get", middlewareChain.Then(http.HandlerFunc(createGetShopByIDHandler(env))))
	http.Handle("/api/v1/shops/update", middlewareChain.Then(http.HandlerFunc(createUpdateShopHandler(env))))

	http.Handle("/api/v1/categories/tree", middlewareChain.Then(http.HandlerFunc(createGetGoodsCategoriesTreeHandler(env))))

	http.Handle("/message", middlewareChain.Then(http.HandlerFunc(createTestMessageHandler(env, hub))))

	http.ListenAndServe(":8080", nil)

}

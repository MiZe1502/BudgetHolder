package handlers

import (
	conf "../configuration"
	env "../env"
	wshub "../hub"
	repos "../repositories"
	"encoding/json"
	"fmt"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
)

// Config stores logger configuration
type Config struct {
	IsHTTPS  bool   `json:"https"`
	CertPath string `json:"certPath"`
	KeyPath  string `json:"keyPath"`
	Port     int    `json:"port"`
}

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

// parseLoggerConfig parses json config for logger
func parseLoggerConfig(config []byte) (Config, error) {
	var parsedConfig Config

	err := json.Unmarshal(config, &parsedConfig)

	return parsedConfig, err
}

func formStringPort(port int) string {
	return ":" + fmt.Sprint(port)
}

// InitHandlers initialize all endpoints with middlewares and handlers
func InitHandlers(env *env.Env, hub *wshub.Hub) {
	config, err := conf.ReadServerConfig(conf.EnvironmentKey(env.EnvironmentKey))
	if err != nil {
		env.Logger.Fatal("Error reading server config: " + err.Error())
	}

	parsedConfig, err := parseLoggerConfig(config)
	if err != nil {
		env.Logger.Fatal("Error parsing server config: " + err.Error())
	}

	middlewareChain := alice.New(createMiddleware(env, Trace),
		createMiddleware(env, Logger),
		createMiddleware(env, CheckSession))

	http.Handle("/", middlewareChain.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/ws", middlewareChain.Then(http.HandlerFunc(createWsHandler(env, hub))))

	//TODO: configure cors

	http.Handle("/api/v1/user/auth", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createAuthHandler(env)))))
	http.Handle("/api/v1/user/new", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createNewUserHandler(env)))))
	http.Handle("/api/v1/user/sync", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createActualizeUserLastOnlineHandler(env)))))
	http.Handle("/api/v1/user/full", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetFullUserInfoHandler(env)))))
	http.Handle("/api/v1/user/group/new", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createNewUserGroupHandler(env)))))
	http.Handle("/api/v1/user/logout", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createCloseUserSessionHandler(env)))))
	http.Handle("/api/v1/user/update", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createUpdateUserHandler(env)))))

	http.Handle("/api/v1/shops/slice", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetShopsSliceHandler(env)))))
	http.Handle("/api/v1/shops/top", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetTopShopsByNameHandler(env)))))
	http.Handle("/api/v1/shops/new", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createAddNewShopHandler(env)))))
	http.Handle("/api/v1/shops/remove", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createRemoveShopHandler(env)))))
	http.Handle("/api/v1/shops/get", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetShopByIDHandler(env)))))
	http.Handle("/api/v1/shops/update", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createUpdateShopHandler(env)))))

	http.Handle("/api/v1/categories/tree", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetGoodsCategoriesTreeHandler(env)))))
	http.Handle("/api/v1/categories/chain", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetCategoryChainByParentIDHandler(env)))))
	http.Handle("/api/v1/categories/category", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetSingleCategoryByIDHandler(env)))))
	http.Handle("/api/v1/categories/list", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetSimpleCategoriesListHandler(env)))))
	http.Handle("/api/v1/categories/remove", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createRemoveCategoryHandler(env)))))
	http.Handle("/api/v1/categories/new", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createAddNewCategoryHandler(env)))))
	http.Handle("/api/v1/categories/update", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createUpdateCategoryHandler(env)))))

	http.Handle("/api/v1/purchases/slice", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetPurchasesWithGoodsDataSliceHandler(env)))))
	http.Handle("/api/v1/purchases/remove", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createRemovePurchaseWithGoodsDetailsHandler(env)))))
	http.Handle("/api/v1/purchases/details/remove", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createRemoveGoodsDetailsItemHandler(env)))))
	http.Handle("/api/v1/purchases/details/update", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createUpdateGoodsDetailsItemHandler(env)))))
	http.Handle("/api/v1/purchases/new", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createAddNewPurchaseWithGoodsDataHandler(env)))))
	http.Handle("/api/v1/purchases/update", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createUpdatePurchaseHandler(env)))))

	http.Handle("/api/v1/goods/slice", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetGoodsSliceHandler(env)))))
	http.Handle("/api/v1/goods/get", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetGoodsItemByIDHandler(env)))))
	http.Handle("/api/v1/goods/top", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createGetTopGoodsItemByNameHandler(env)))))
	http.Handle("/api/v1/goods/remove", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createRemoveGoodsItemHandler(env)))))
	http.Handle("/api/v1/goods/new", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createAddNewGoodsItemHandler(env)))))
	http.Handle("/api/v1/goods/update", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createUpdateGoodsItemHandler(env)))))

	http.Handle("/message", cors.AllowAll().Handler(middlewareChain.Then(http.HandlerFunc(createTestMessageHandler(env, hub)))))

	if parsedConfig.IsHTTPS {
		http.ListenAndServeTLS(formStringPort(parsedConfig.Port), parsedConfig.CertPath, parsedConfig.KeyPath, nil)
	} else {
		http.ListenAndServe(formStringPort(parsedConfig.Port), nil)
	}
}

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
	FrontEnd string `json:"frontEnd"`
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

func configureCors(config Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{config.FrontEnd},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
}

func createHandlerWrappper(mdc alice.Chain, c *cors.Cors) func(func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return func(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
		return c.Handler(mdc.Then(http.HandlerFunc(handler)))
	}
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

	//middlewareChain
	mdc := alice.New(createMiddleware(env, Trace),
		createMiddleware(env, Logger),
		createMiddleware(env, CheckSession))

	http.Handle("/", mdc.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/ws", mdc.Then(http.HandlerFunc(createWsHandler(env, hub))))

	c := configureCors(parsedConfig)
	wrap := createHandlerWrappper(mdc, c)

	http.Handle("/api/v1/user/auth", wrap(createAuthHandler(env)))
	http.Handle("/api/v1/user/new", wrap(createNewUserHandler(env)))
	http.Handle("/api/v1/user/full", wrap(createGetFullUserInfoHandler(env)))
	http.Handle("/api/v1/user/group/new", wrap(createNewUserGroupHandler(env)))
	http.Handle("/api/v1/user/logout", wrap(createCloseUserSessionHandler(env)))
	http.Handle("/api/v1/user/update", wrap(createUpdateUserHandler(env)))

	http.Handle("/api/v1/shops/slice", wrap(createGetShopsSliceHandler(env)))
	http.Handle("/api/v1/shops/list", wrap(createGetSimpleShopsListHandler(env)))
	http.Handle("/api/v1/shops/new", wrap(createAddNewShopHandler(env)))
	http.Handle("/api/v1/shops/remove", wrap(createRemoveShopHandler(env)))
	http.Handle("/api/v1/shops/get", wrap(createGetShopByIDHandler(env)))
	http.Handle("/api/v1/shops/update", wrap(createUpdateShopHandler(env)))

	http.Handle("/api/v1/categories/tree", wrap(createGetGoodsCategoriesTreeHandler(env)))
	http.Handle("/api/v1/categories/chain", wrap(createGetCategoryChainByParentIDHandler(env)))
	http.Handle("/api/v1/categories/category", wrap(createGetSingleCategoryByIDHandler(env)))
	http.Handle("/api/v1/categories/list", wrap(createGetSimpleCategoriesListHandler(env)))
	http.Handle("/api/v1/categories/remove", wrap(createRemoveCategoryHandler(env)))
	http.Handle("/api/v1/categories/new", wrap(createAddNewCategoryHandler(env)))
	http.Handle("/api/v1/categories/update", wrap(createUpdateCategoryHandler(env)))

	http.Handle("/api/v1/purchases/slice", wrap(createGetPurchasesWithGoodsDataSliceHandler(env)))
	http.Handle("/api/v1/purchases/remove", wrap(createRemovePurchaseWithGoodsDetailsHandler(env)))
	http.Handle("/api/v1/purchases/details/remove", wrap(createRemoveGoodsDetailsItemHandler(env)))
	http.Handle("/api/v1/purchases/details/update", wrap(createUpdateGoodsDetailsItemHandler(env)))
	http.Handle("/api/v1/purchases/new", wrap(createAddNewPurchaseWithGoodsDataHandler(env)))
	http.Handle("/api/v1/purchases/update", wrap(createUpdatePurchaseHandler(env)))

	http.Handle("/api/v1/goods/slice", wrap(createGetGoodsSliceHandler(env)))
	http.Handle("/api/v1/goods/get", wrap(createGetGoodsItemByIDHandler(env)))
	http.Handle("/api/v1/goods/list", wrap(createGetSimpleGoodsItemsListHandler(env)))
	http.Handle("/api/v1/goods/remove", wrap(createRemoveGoodsItemHandler(env)))
	http.Handle("/api/v1/goods/new", wrap(createAddNewGoodsItemHandler(env)))
	http.Handle("/api/v1/goods/update", wrap(createUpdateGoodsItemHandler(env)))

	http.Handle("/message", wrap(createTestMessageHandler(env, hub)))

	if parsedConfig.IsHTTPS {
		http.ListenAndServeTLS(formStringPort(parsedConfig.Port), parsedConfig.CertPath, parsedConfig.KeyPath, nil)
	} else {
		http.ListenAndServe(formStringPort(parsedConfig.Port), nil)
	}
}

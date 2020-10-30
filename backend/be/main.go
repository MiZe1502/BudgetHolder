package main

import (
	"fmt"
	"os"

	db "./db"
	repos "./repositories"

	utils "./utils"
)

func main() {
	log, err := utils.InitLogger("Dev")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := db.Connect("Dev")

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	hub := createHub()

	tokenGenerator, err := utils.InitTokenGenerator("Dev")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	context := &Env{db: conn, logger: log, hub: hub, token: tokenGenerator}

	go context.hub.runHub(context)

	defer conn.Close()
	defer hub.closeAllConnections(context)

	// uuid := utils.GetNewUUID()

	// fmt.Println(uuid)

	// uuidParsed, err := utils.GetUUIDfromString(uuid)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// token, err := context.token.CreateNewToken(uuidParsed)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(token)

	// tkn, err := context.token.ParseToken(token)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(tkn.SessionID)

	//testClosure(context)()
	initHandlers(context)
}

func testClosure(env *Env) func() {
	return func() {
		var repo repos.ShopsRepository

		repo.SetDb(env.db)

		shops, err := repo.GetSlice(0, 5)
		if err != nil {
			env.logger.Error(err.Error())
		}

		for i := 0; i < len(shops); i++ {
			env.logger.Info(fmt.Sprint(shops[i].ID))
		}

		shop, err := repo.GetEntityByID(3)
		if err != nil {
			env.logger.Error(err.Error())
		}

		env.logger.Info(shop.Name)

		shopID, err := repo.RemoveEntityByID(2, utils.GetNewUUID())

		if err != nil {
			env.logger.Error(err.Error())
		}

		env.logger.Info(fmt.Sprint(shopID))

	}
}

// func getHandlers(env *Env) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//get data from db
// 	}
// }

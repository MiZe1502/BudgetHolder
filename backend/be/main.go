package main

import (
	"fmt"
	"os"

	db "./db"
	env "./env"
	handlers "./handlers"
	hub "./hub"

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

	hub := hub.CreateHub()

	tokenGenerator, err := utils.InitTokenGenerator("Dev")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	context := &env.Env{EnvironmentKey: "Dev",
		Db:     conn,
		Logger: log,
		Token:  tokenGenerator}

	go hub.RunHub(context)
	go log.ProcessLogFileUpdate()

	defer conn.Close()
	defer hub.CloseAllConnections(context)

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
	handlers.InitHandlers(context, hub)
}

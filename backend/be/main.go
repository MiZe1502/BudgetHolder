package main

import (
	"fmt"

	db "./db"
	repos "./repositories"
)

func main() {
	conn := db.Connect("Dev")

	context := &Env{db: conn}

	defer conn.Close()

	testClosure(context)()
}

func testClosure(env *Env) func() {
	return func() {

		var repo repos.ShopsRepository

		repo.SetDb(env.db)

		shops, err := repo.GetSlice(0, 5)

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(shops); i++ {
			fmt.Println(shops[i].ID)
		}

		shop, err := repo.GetEntityByID(3)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(shop.Name)

		shopID, err := repo.RemoveEntityByID(2, "124e4567-e89b-12d3-a456-426655440000")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(shopID)

	}
}

// func getHandlers(env *Env) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//get data from db
// 	}
// }

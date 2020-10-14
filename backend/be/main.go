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

		shops, err := repo.GetSlice(env.db, 0, 5)

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(shops); i++ {
			fmt.Println(shops[i].ID)
		}
	}
}

// func getHandlers(env *Env) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//get data from db
// 	}
// }

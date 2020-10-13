package main

import (
	db "./db"
	"github.com/georgysavva/scany/pgxscan"

	"context"
	"fmt"
)

type Shop struct {
	Id      int
	Name    string
	Address string
	Comment string
	Url     string
}

func main() {
	conn := db.Connect("Dev")

	context := &Env{db: conn}

	defer conn.Close()

	testClosure(context)()
}

func testClosure(env *Env) func() {
	return func() {
		var shops []*Shop

		err := pgxscan.Select(context.Background(), env.db, &shops, `SELECT id, name, comment from budget.get_shops(0, 5)`)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(shops)
	}
}


// func getHandlers(env *Env) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//get data from db
// 	}
// }

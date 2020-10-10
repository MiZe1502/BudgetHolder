package main

import (
	db "./db"
	"github.com/georgysavva/scany/pgxscan"

	"context"
	"fmt"
	// "github.com/georgysavva/scany/pgxscan"
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

	defer conn.Close()

	var shops []*Shop

	// rows, err := conn.Query(context.Background(), "SELECT id, name, comment from budget.shops")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var shop *Shop = {}
	// 	err = rows.Scan(&shop.Id, &shop.Name)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(shop)

	// }
	err := pgxscan.Select(context.Background(), conn, &shops, `SELECT id, name, comment from budget.get_shops(0, 5)`)
	//err := pgxscan.Select(context.Background(), conn, &shops, `SELECT id, name, comment from budget.shops`)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(shops)
}

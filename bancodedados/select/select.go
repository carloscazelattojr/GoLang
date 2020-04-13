package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	for rows.Next() {
		var users usuario
		rows.Scan(&users.id, &users.nome)
		fmt.Println(users)
	}

}

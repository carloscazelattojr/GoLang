package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	transacao, _ := db.Begin()
	stmt, _ := transacao.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(20, "Bia")
	stmt.Exec(21, "Eliana")
	stmt.Exec(22, "Miti")

	_, err = stmt.Exec(1, "Carlos") //Forcando dar error de duplicidade. (o id: 1 já está cadastrado.)
	if err != nil {
		transacao.Rollback()
		log.Fatal(err)
	}

	transacao.Commit()
}

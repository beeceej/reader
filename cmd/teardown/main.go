package main

import (
	r "github.com/beeceej/reader"
	"github.com/beeceej/reader/db"
	"github.com/jmoiron/sqlx"
)

func main() {
	r.NewReader(map[r.EnvKey]r.EnvVal{}).
		Bind(db.MySQLReader).
		With(teardown)
}

func teardown(env r.Env) {
	conn := db.GetMySQLConnection.Run(env).(*sqlx.DB)
	dropTable(conn)
}

func dropTable(conn *sqlx.DB) {
	query := `
DROP TABLE puppies;
`
	_, err := conn.Exec(query)

	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	r "github.com/beeceej/reader"
	"github.com/beeceej/reader/db"
	"github.com/jmoiron/sqlx"
)

func main() {
	r.NewReader(map[r.EnvKey]r.EnvVal{}).
		Bind(db.MySQLReader).
		With(setup)
}

func setup(env r.Env) {
	conn := db.GetMySQLConnection.Run(env).(*sqlx.DB)
	createTable(conn)
	insertData(conn)
}

func createTable(conn *sqlx.DB) {
	sql := `CREATE TABLE IF NOT EXISTS puppies (
    id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    breed VARCHAR(255) NOT NULL ,
    color VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);`

	_, err := conn.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func insertData(conn *sqlx.DB) {
	insertSQL := `
INSERT INTO puppies
    (name, breed, color)
		VALUES
		('Kona', 'Mini Aussie', 'Red Tri'),
		('Yogi', 'Doodle', 'Golden'),
		('Hiro', 'Shiba', 'Brown'),
		('Watons', 'Shiba', 'Black');
`

	_, err := conn.Exec(insertSQL)
	if err != nil {
		panic(err.Error())
	}
}

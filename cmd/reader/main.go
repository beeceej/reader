package main

import (
	"fmt"
	"os"

	r "github.com/beeceej/reader"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	keyMySQLHost             = "keyMySQLHost"
	keyMySQLPort             = "keyMySQLPort"
	keyMySQLDB               = "keyMySQLDB"
	keyMySQLUser             = "keyMySQLUser"
	keyMySQLPassword         = "keyMySQLPassword"
	keyMySQLConnectionString = "keyMySQLConnectionString"
	keyMySQLConnection       = "keyMySQLConnection"
)

var mysqlHost = func(env r.Env) r.MonadReader {
	return r.MonadReader{
		r.KVReader(env, keyMySQLHost, os.Getenv("MYSQL_HOST")),
		env,
	}
}

var mysqlPort = func(env r.Env) r.MonadReader {
	return r.MonadReader{
		r.KVReader(env, keyMySQLPort, os.Getenv("MYSQL_PORT")),
		env,
	}
}

var mysqlDB = func(env r.Env) r.MonadReader {
	return r.MonadReader{
		r.KVReader(env, keyMySQLDB, os.Getenv("MYSQL_NAME")),
		env,
	}
}

var mysqlUser = func(env r.Env) r.MonadReader {
	return r.MonadReader{
		r.KVReader(env, keyMySQLUser, os.Getenv("MYSQL_USER")),
		env,
	}
}

var mysqlPassword = func(env r.Env) r.MonadReader {
	return r.MonadReader{
		r.KVReader(env, keyMySQLPassword, os.Getenv("MYSQL_PASS")),
		env,
	}
}

var mysqlConnectionString = func(env r.Env) r.MonadReader {
	host := getMySQLHost.Run(env).(string)
	port := getMySQLPort.Run(env).(string)
	db := getMySQLDB.Run(env).(string)
	user := getMySQLUser.Run(env).(string)
	password := getMySQLPassword.Run(env).(string)

	if password != "" {
		password = ":" + password
	}

	connDetails := fmt.Sprintf("%s%s@(%s:%s)/%s?parseTime=true", user, password, host, port, db)

	return r.MonadReader{
		r.KVReader(env, keyMySQLConnectionString, connDetails),
		env,
	}
}

var mysqlConnection = func(env r.Env) r.MonadReader {
	connString := getMySQLConnectionString.Run(env).(string)
	db, err := sqlx.Connect("mysql", connString)

	if err != nil {
		panic(err.Error())
	}

	return r.MonadReader{
		r.KVReader(env, keyMySQLConnection, db),
		env,
	}
}

var getMySQLHost = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLHost]
})

var getMySQLPort = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLPort]
})

var getMySQLDB = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLDB]
})

var getMySQLUser = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLUser]
})

var getMySQLPassword = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLPassword]
})

var getMySQLConnectionString = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLConnectionString]
})

var getMySQLConnection = r.AReader(func(env r.Env) r.EnvVal {
	return env[keyMySQLConnection]
})

func main() {
	envReader := r.NewReader(map[r.EnvKey]r.EnvVal{}).
		Bind(mysqlHost).
		Bind(mysqlPort).
		Bind(mysqlDB).
		Bind(mysqlUser).
		Bind(mysqlPassword).
		Bind(mysqlConnectionString).
		Bind(mysqlConnection)

	envReader.With(doDatabaseStuff)
}

func doDatabaseStuff(env r.Env) {
	db := getMySQLConnection.Run(env).(*sqlx.DB)
	createTable(db)
	insertData(db)
	queryData(db)
	dropTable(db)
}

func createTable(db *sqlx.DB) {
	sql := `CREATE TABLE IF NOT EXISTS puppies (
    id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    breed VARCHAR(255) NOT NULL ,
    color VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func insertData(db *sqlx.DB) {
	insertSQL := `
INSERT INTO puppies
    (name, breed, color)
		VALUES
		('Kona', 'Mini Aussie', 'Red Tri'),
		('Yogi', 'Doodle', 'Golden'),
		('Hiro', 'Shiba', 'Brown'),
		('Watons', 'Shiba', 'Black');
`

	_, err := db.Exec(insertSQL)
	if err != nil {
		panic(err.Error())
	}

}

func queryData(db *sqlx.DB) {
	query := `
SELECT * FROM puppies;
`

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err = rows.Close(); err != nil {
			panic(err.Error())
		}
	}()

	for rows.Next() {
		var (
			id int

			name, breed, color string
		)
		err = rows.Scan(&id, &name, &breed, &color)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(id, name, breed, color)

	}
}

func dropTable(db *sqlx.DB) {
	query := `
DROP TABLE puppies;
`
	_, err := db.Exec(query)

	if err != nil {
		panic(err.Error())
	}
}

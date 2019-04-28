package db

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
	KeyMySQLConnection       = "keyMySQLConnection"
)

func MySQLReader(env r.Env) r.MonadReader {
	return r.NewReader(env).
		Bind(mysqlHost).
		Bind(mysqlPort).
		Bind(mysqlDB).
		Bind(mysqlUser).
		Bind(mysqlPassword).
		Bind(mysqlConnectionString).
		Bind(mysqlConnection)
}

func MySQL(env r.Env) *sqlx.DB {
	conn := GetMySQLConnection.Run(env).(*sqlx.DB)
	return conn
}

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
		r.KVReader(env, KeyMySQLConnection, db),
		env,
	}
}

var GetMySQLConnection = r.AReader(func(env r.Env) r.EnvVal {
	return env[KeyMySQLConnection]
})

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

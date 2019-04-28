package puppies

import (
	r "github.com/beeceej/reader"
	"github.com/beeceej/reader/db"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	r.MonadReader
	db *sqlx.DB
}

const (
	KeyRepository = "keyPuppiesRepository"
)

func RepositoryReader(env r.Env) r.MonadReader {
	return r.KVReader(
		env,
		KeyRepository,
		Repository{
			db: env[db.KeyMySQLConnection].(*sqlx.DB),
		},
	)
}

var GetRepository = r.AReader(func(env r.Env) r.EnvVal {
	return env[KeyRepository]
})

func (p Repository) GetPuppy(wantedID int) (Puppy, error) {
	query := `SELECT * FROM puppies where id = ?;`
	rows, err := p.db.Queryx(query, wantedID)

	if err != nil {
		return Puppy{}, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			panic(err.Error())
		}
	}()
	var puppy Puppy
	for rows.Next() {
		rows.StructScan(&puppy)
	}

	return puppy, nil
}

func (p Repository) InsertPuppy(puppy Puppy) error {
	sql := `
INSERT INTO puppies
  (name, breed, color)
VALUES
  (?, ?, ?);`
	_, err := p.db.Exec(sql, puppy.Name, puppy.Breed, puppy.Color)
	return err
}

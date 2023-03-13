package psql

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/vmoltaemcrkonrgcechd/pocu/config"
)

type PSQL struct {
	DB *sql.DB
	Sq sq.StatementBuilderType
}

func New(cfg *config.Config) (*PSQL, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PSQL{
		DB: db,
		Sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db),
	}, nil
}

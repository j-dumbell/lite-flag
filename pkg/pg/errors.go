package pg

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var ErrAlreadyExists = errors.New("record already exists")

var ErrNoRows = errors.New("no rows in result set")

func ParseError(err error) error {
	if err == nil {
		return nil
	}

	pgErr, ok := err.(*pq.Error)
	if ok {
		if pgErr.Code == "23505" {
			return ErrAlreadyExists
		}
	}

	if err == sql.ErrNoRows {
		return ErrNoRows
	}

	return err
}

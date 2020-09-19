package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func UserPut(db *sql.DB, id int, name string) error {
	if _, err := db.Exec("INSERT INTO users(id, Name) Values ($1, $2)", id, name); err != nil {
		return err
	}

	return nil
}

func TestUserPut(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO users").WithArgs(1, "Bond").WillReturnResult(sqlmock.NewResult(1, 1))

	err = UserPut(db, 1, "Bond")
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

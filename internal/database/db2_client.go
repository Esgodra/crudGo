package database

import (
	"database/sql"
	"fmt"
	_ "github.com/ibmdb/go_ibm_db" //Libreria de db2"
)

type DB2Client struct {
	*sql.DB
}

func NewDb2Client(source string) *DB2Client {
	db, err := sql.Open("go_ibm_db", source)

	if err != nil {
		_ = fmt.Errorf("cannot create db tentat: %s", err.Error())
		panic("..")
	}

	err = db.Ping()

	if err != nil {
		fmt.Printf("Cannot connect to db2")
	}

	return &DB2Client{DB: db}

}

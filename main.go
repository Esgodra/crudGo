package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ibmdb/go_ibm_db" //Libreria de db2
)

// Database estructura
type Database struct {
	db     *sql.DB
	schema string
}

var database Database

func main() {
	fmt.Println("Started")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	if isConnect, err := ReconnectDB(); err != nil && !isConnect {
		fmt.Println("Error")
	} else {
		fmt.Println("Connected")
	}

	log.Fatal(http.ListenAndServe(":444", nil))
}

//ReconnectDB function for connect to Database
func ReconnectDB() (bool, error) {
	dsn := "DATABASE=testdb;HOSTNAME=172.17.0.1;UID=db2inst1;PWD=IBM2BLU;PORT=50000;PROTOCOL=TCPIP;"
	if isConnect, err := OpenDBConnection(dsn); err != nil && !isConnect {
		fmt.Printf("Error, while connecting to DB: %s \n", err)
		return false, err
	}
	return true, nil
}

// OpenDBConnection function open this database connection
func OpenDBConnection(dsn string) (bool, error) {
	var err error
	database.db, err = sql.Open("go_ibm_db", dsn)
	if err != nil {
		return false, err
	}
	return true, nil
}

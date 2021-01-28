package main

import (
	"crudGo/internal/database"
	"crudGo/server"
	"crudGo/web"
	"database/sql"
	"fmt"
)

func execquery(st *sql.Stmt) error {
	rows, err := st.Query()
	if err != nil {
		return err
	}
	cols, _ := rows.Columns()
	fmt.Printf("%s  %s  %s  %s\n", cols[0], cols[1], cols[2], cols[3])
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var t, x, m, n string
		err = rows.Scan(&t, &x, &m, &n)
		if err != nil {
			return err
		}
		fmt.Printf("%v %v %v       %v\n", t, x, m, n)
	}
	return nil
}

func getBooks(db *sql.DB) error {
	st, err := db.Prepare("SELECT * FROM BOOKS")
	if err != nil {
		return err
	}
	err = execquery(st)
	if err != nil {
		return err
	}
	return nil
}

func updateBook(bookID int, title string, author string, readstatus int, db *sql.DB) error {
	st, err := db.Prepare("UPDATE BOOKS SET TITLE = ?, AUTHOR= ?, READSTATUS= ? WHERE ID = ?")

	if err != nil {
		return err
	}
	err = execquery(st)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Main")
	client := database.NewDb2Client("DATABASE=BLUDB;HOSTNAME=dashdb-txn-sbox-yp-dal09-10.services.dal.bluemix.net;PORT=50000;PROTOCOL=TCPIP;UID=wpd52601;PWD=ccj779@zvp40hz2d;")

	handler := web.NewCreateBookHandler(client)
	mux := server.Routes(handler)
	serv := server.NewServer(mux)
	serv.Run()
}

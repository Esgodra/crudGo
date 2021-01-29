package main

import (
	"crudGo/internal/database"
	"crudGo/server"
	"crudGo/web"
	"fmt"
)

func main() {
	fmt.Println("Main")
	client := database.NewDb2Client("DATABASE=BLUDB;HOSTNAME=dashdb-txn-sbox-yp-dal09-10.services.dal.bluemix.net;PORT=50000;PROTOCOL=TCPIP;UID=wpd52601;PWD=ccj779@zvp40hz2d;")

	create := web.NewCreateBookHandler(client)
	update := web.NewUpdateBookHandler(client)
	delete := web.NewDeleteBookHandler(client)
	get := web.NewGetBookHandler(client)
	mux := server.Routes(create, update, delete, get)
	serv := server.NewServer(mux)
	serv.Run()
}

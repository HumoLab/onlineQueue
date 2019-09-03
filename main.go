package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/onlineQueue/backend/database"
)

func main() {
	//==========   DB  ==========
	err := database.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer database.Disconnect()
	//===========================
}

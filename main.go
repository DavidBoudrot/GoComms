package main

// main method that will start http server, start db, form db connection

import (
	"fmt"
	"messenger/utils"
)

func main() {
	db, err := utils.Connect()
	if err != nil {
		fmt.Print("Error connecting to database: ", err)
	}
	fmt.Print(db)
	utils.HTTPStartup()
}

package main

// main method that will start http server, start db, form db connection

import (
	"messenger/utils"
)

func main() {
	utils.Connect()
	utils.HTTPStartup()
}

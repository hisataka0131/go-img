package main

import (
	"go-img/db"
	"go-img/router"
)

func main() {
	// DB接続
	db.InitDb()
	defer db.Close()

	router.Router()
}

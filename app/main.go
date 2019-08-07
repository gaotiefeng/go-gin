package main

import (
	"go-gin/app/router"
	db "go-gin/database"
)

func main() {

	defer db.SqlDB.Close()

	router := router.SetupRouter()

	_ = router.Run()
}

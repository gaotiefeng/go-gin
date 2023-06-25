package main

import (
	"go-gin/app/router"
)

func main() {

	router := router.SetupRouter()
	_ = router.Run()
}

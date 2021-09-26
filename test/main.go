package main

import (
	"test/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run(":8080")
}

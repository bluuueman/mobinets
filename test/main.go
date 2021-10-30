package main

import (
	"test/router"
	"test/utility"
)

//*************************************************************8
/*
func main() {

	var output bytes.Buffer
	if utility.DockerImages(&output) {

		log.Println(output.String())
	}
}

//*************************************************************8
*/

func main() {
	if !utility.InitCarControl() {
		return
	}
	router := router.SetupRouter()
	_ = router.Run(":8080")
	utility.CloesCarControl()
	return
}

package main

import (
	"test/router"
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
	router := router.SetupRouter()
	_ = router.Run(":8080")
}

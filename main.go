package main

import (
	"go-rest-mongo/internal/app/api"
)

func main() {

	router := api.Init()

	router.Run(":8001")

}

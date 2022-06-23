package main

import (
	model "server/models"
	r "server/routers"
)

func main() {
	model.ConnectDB()

	r.InitRouter()
}

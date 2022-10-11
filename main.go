package main

import (
	"tfalc/Api-gin/database"
	"tfalc/Api-gin/routes"
)

func main() {
	database.ConectaComBanco()
	routes.HandleRequests()
}

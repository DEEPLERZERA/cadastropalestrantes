package main

import (
	"cadastropalestrantes/database"
	"cadastropalestrantes/routes"
)


func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
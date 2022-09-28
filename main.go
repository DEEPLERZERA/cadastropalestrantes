package main

import (
	"CadastroPalestrantes/database"
	"CadastroPalestrantes/routes"
)


func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
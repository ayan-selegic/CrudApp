package main

import (
	a "CrudSystem/config"
	b "CrudSystem/routes"
)

func main() {
	a.InitialMigration() // establish DB connection
	b.InitializeRouter() // creating and initializing routers with MUX
}

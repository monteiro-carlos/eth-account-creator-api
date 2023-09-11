package main

import (
	"eth-account-creator-api/internal/container"
	"eth-account-creator-api/internal/routes"
)

// @title Ethereum Account API
// @version 1.0
// @description This API is used to create Ethereum Accounts and make transaction using ETH.
// @termsOfService http://swagger.io/terms/
// @contact.name Carlos Fernandes
// @query.collection.format multi
// @in header
// @schemes http https.
func main() {
	dep, err := container.New()
	if err != nil {
		panic(err)
	}
	routes.Handler(dep)
}

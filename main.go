package main

import (
	"log"
	routes "github.com/chocogem/keycloak-authen-golang-clean-arch/deliveries/routes"
)

func main() {

	server, err := routes.SetupNewRouter()
	if err != nil {
		log.Fatalf("Internal server error: %v", err)
	}

	server.Run(":8081")
}

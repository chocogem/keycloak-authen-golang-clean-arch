package routes

import (
	"github.com/chocogem/keycloak-authen-golang-clean-arch/databases"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/deliveries"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/repositories"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/usecases"
	"github.com/gin-gonic/gin"
)

func SetupNewRouter() (*gin.Engine, error) {
	router := gin.Default()

	repo := repositories.NewTestRepository(databases.GetMockData())
	usecase := usecases.NewTestUseCase(repo)
	delivery := deliveries.NewTestDelivery(usecase)
	routeGroup := router.Group("/v1")

	routeGroup.Use(usecases.AuthenMiddleware())
	{
		routeGroup.GET("test/:id", delivery.Get)
	}

	return router, nil
}


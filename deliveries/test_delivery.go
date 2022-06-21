package deliveries

import (
	"net/http"
	"strconv"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/domains"
	"github.com/gin-gonic/gin"
)

type testDeliveries struct {
	testUseCase domains.ITestRepository
}

func NewTestDelivery(useCase domains.ITestRepository) *testDeliveries {
	return &testDeliveries{useCase}
}

func (d *testDeliveries) Get(c *gin.Context) {
	idStr := c.Param("id")
	if len(idStr) >0{
		id, err := strconv.Atoi(idStr)
		if id > 0 && err == nil {
			result, err := d.testUseCase.Get(id)
			if err == nil {
				 c.JSON(http.StatusOK, result)
				 return
			}
			 c.AbortWithStatus(http.StatusNotFound)
			 return
		}
	}
	 c.AbortWithStatus(http.StatusBadGateway)
	 return
	
}

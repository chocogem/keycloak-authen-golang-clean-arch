package databases

import (
	"github.com/chocogem/keycloak-authen-golang-clean-arch/models"
)

type MockDatabase struct {
	Mock map[int]models.Test
}

func GetMockData() MockDatabase {
	mockMap := make(map[int]models.Test)
	mockMap[10001] = models.Test{
		ID:        10001,
		FirstName: "TestName",
		LastName:  "TestLastName",
	}
	return MockDatabase{mockMap}
}

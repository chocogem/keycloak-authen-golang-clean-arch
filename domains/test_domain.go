package domains

import "github.com/chocogem/keycloak-authen-golang-clean-arch/models"

type ITestRepository interface {
	Get(id int) (models.Test, error)
}


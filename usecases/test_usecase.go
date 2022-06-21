package usecases

import (
	"github.com/chocogem/keycloak-authen-golang-clean-arch/domains"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/models"
)

type testUseCase struct {
	Repo domains.ITestRepository
}

func NewTestUseCase(repo domains.ITestRepository) domains.ITestRepository {
	return &testUseCase{repo}
}

func (uc *testUseCase) Get(id int) (models.Test, error) {

	return uc.Repo.Get(id)
}

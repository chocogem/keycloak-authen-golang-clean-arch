package repositories

import (
	"github.com/chocogem/keycloak-authen-golang-clean-arch/databases"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/domains"
	"github.com/chocogem/keycloak-authen-golang-clean-arch/models"
)

type TestRepository struct{
	data databases.MockDatabase
}

func NewTestRepository(db databases.MockDatabase) domains.ITestRepository {
	return &TestRepository{db}
}

func (repo *TestRepository) Get(id int) (models.Test, error) {
	return repo.data.Mock[id], nil
}


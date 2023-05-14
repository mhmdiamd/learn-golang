package repository

import (
	"golang-unit-test/8-mock-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategortyRepository interface {
	FindById(id string) *entity.Category
}

type CategortyRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategortyRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entity.Category)
	return &category
}

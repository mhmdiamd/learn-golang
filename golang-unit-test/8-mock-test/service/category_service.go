package service

import (
	"errors"
	"golang-unit-test/8-mock-test/entity"
	"golang-unit-test/8-mock-test/repository"
)

type CategoryService struct {
	Repository repository.CategortyRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}
}

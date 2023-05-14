package service

import (
	"golang-unit-test/8-mock-test/entity"
	"golang-unit-test/8-mock-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategortyRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{
	Repository: categoryRepository,
}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// Get cateogry
	cateogory, err := categoryService.Get("1")

	assert.NotNil(t, err)
	assert.Nil(t, cateogory)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	// Melakukan Programm dengan mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}

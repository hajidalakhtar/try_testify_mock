package service

import (
	"belajar_golang/entity"
	"belajar_golang/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {

	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, error := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, error)
}

func TestCategoryService_Get2(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, error := categoryService.Get("2")
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}

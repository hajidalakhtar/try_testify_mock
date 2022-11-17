package service

import (
	"belajar_golang/entity"
	"belajar_golang/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not found")
	} else {
		return category, nil
	}
}

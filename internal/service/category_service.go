package service

import (
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/database"
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDb database.CategoryDB) *CategoryService{
	return &CategoryService{CategoryDB: categoryDb}
}

func (service *CategoryService) GetCategories()([]*entity.Category, error){
	categories, err := service.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (service *CategoryService) CreateCategory(name string)(*entity.Category, error){
	category := entity.NewCategory(name)
	_, err := service.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (serice *CategoryService) GetCategory(id string)(*entity.Category, error){
	category, err := serice.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
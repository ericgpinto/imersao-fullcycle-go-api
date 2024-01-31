package service

import (
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/database"
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService{
	return &ProductService{ProductDB: productDB}
}

func (service *ProductService) GetProducts()([]*entity.Product, error){
	products, err := service.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *ProductService) GetProcuct(id string)(*entity.Product, error){
	product, err := service.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (service *ProductService) GetProcuctByCategoryID(categoryID string)([]*entity.Product, error){
	products, err := service.ProductDB.GetProductByCategoryId(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}


func (service *ProductService) CreateProduct(name, description, category_id, image_url string, price float64)(*entity.Product, error){
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := service.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product,nil
}

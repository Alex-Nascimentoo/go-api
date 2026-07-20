package usecase

import (
	"github.com/alex-nascimentoo/go-api/model"
	"github.com/alex-nascimentoo/go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{repository: repo}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	return pu.repository.GetProductById(id)
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	prodId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = prodId

	return product, nil
}

package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ly/internal/repositories"
	"ly/pkg/models"
)

type ProductTypeService interface {
	GetAll(ctx context.Context, name string) ([]*models.ProductType, error)
	GetById(ctx context.Context, id int) (*models.ProductType, error)
	Create(ctx context.Context, productType *models.ProductType) error
	Update(ctx context.Context, productType *models.ProductType) error
	Delete(ctx context.Context, id int) error
}

type productTypeService struct {
	db                    *gorm.DB
	productTypeRepository repositories.ProductTypeRepository
	productService        ProductService
}

func (s *productTypeService) GetAll(ctx context.Context, name string) ([]*models.ProductType, error) {
	return s.productTypeRepository.GetProductTypeList(ctx, s.db, name)
}

func (s *productTypeService) GetById(ctx context.Context, id int) (*models.ProductType, error) {
	return s.productTypeRepository.GetProductTypeById(ctx, s.db, id)
}

func (s *productTypeService) Create(ctx context.Context, productType *models.ProductType) error {
	name, err := s.productTypeRepository.GetProductTypeByName(ctx, s.db, productType.Name)
	if err == nil && name != nil {
		return errors.New("product type already exists")
	}
	return s.productTypeRepository.CreateProduct(ctx, s.db, productType)
}

func (s *productTypeService) Update(ctx context.Context, productType *models.ProductType) error {
	name, err := s.productTypeRepository.GetProductTypeByName(ctx, s.db, productType.Name)
	if err == nil && name != nil {
		return errors.New("product type already exists")
	}
	return s.productTypeRepository.UpdateProductType(ctx, s.db, productType)
}

func (s *productTypeService) Delete(ctx context.Context, id int) error {
	productVos, _ := s.productService.GetProductByProductTypeId(ctx, id)
	if len(productVos) > 0 {
		return errors.New("product type has products, can't delete")
	}
	return s.productTypeRepository.DeleteProductType(ctx, s.db, id)
}

func NewProductTypeService(
	db *gorm.DB,
	productTypeRepository repositories.ProductTypeRepository,
	productService ProductService,
) ProductTypeService {
	return &productTypeService{
		db:                    db,
		productTypeRepository: productTypeRepository,
		productService:        productService,
	}
}

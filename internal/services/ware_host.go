package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ly/internal/repositories"
	"ly/pkg/models"
)

type WareHostService interface {
	GetAll(ctx context.Context, name string) ([]*models.WareHost, error)
	GetById(ctx context.Context, id int) (*models.WareHost, error)
	Create(ctx context.Context, ware *models.WareHost) error
	Update(ctx context.Context, ware *models.WareHost) error
	Delete(ctx context.Context, id int) error
}

type wareHostService struct {
	db                 *gorm.DB
	wareHostRepository repositories.WareHostRepository
	productService     ProductService
}

func (s *wareHostService) GetAll(ctx context.Context, name string) ([]*models.WareHost, error) {
	return s.wareHostRepository.GetProductTypeList(ctx, s.db, name)
}

func (s *wareHostService) GetById(ctx context.Context, id int) (*models.WareHost, error) {
	return s.wareHostRepository.GetProductTypeById(ctx, s.db, id)
}

func (s *wareHostService) Create(ctx context.Context, ware *models.WareHost) error {
	name, err := s.wareHostRepository.GetProductTypeByName(ctx, s.db, ware.Name)
	if err == nil && name != nil {
		return errors.New("product type already exists")
	}
	return s.wareHostRepository.CreateProduct(ctx, s.db, ware)
}

func (s *wareHostService) Update(ctx context.Context, ware *models.WareHost) error {
	name, err := s.wareHostRepository.GetProductTypeByName(ctx, s.db, ware.Name)
	if err == nil && name != nil {
		return errors.New("product type already exists")
	}
	return s.wareHostRepository.UpdateProductType(ctx, s.db, ware)
}

func (s *wareHostService) Delete(ctx context.Context, id int) error {
	productVos, _ := s.productService.GetProductByWareHostId(ctx, id)
	if len(productVos) > 0 {
		return errors.New("product type has products, can't delete")
	}
	return s.wareHostRepository.DeleteProductType(ctx, s.db, id)
}

func NewWareHostService(
	db *gorm.DB,
	wareHostRepository repositories.WareHostRepository,
	productService ProductService,
) WareHostService {
	return &wareHostService{
		db:                 db,
		wareHostRepository: wareHostRepository,
		productService:     productService,
	}
}

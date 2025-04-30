package repositories

import (
	"context"
	"gorm.io/gorm"
	"ly/pkg/models"
)

type ProductTypeRepository interface {
	CreateProduct(ctx context.Context, db *gorm.DB, productType *models.ProductType) error
	GetProductTypeList(ctx context.Context, db *gorm.DB, name string) ([]*models.ProductType, error)
	GetProductTypeById(ctx context.Context, db *gorm.DB, id int) (*models.ProductType, error)
	GetProductTypeByName(ctx context.Context, db *gorm.DB, name string) (*models.ProductType, error)
	UpdateProductType(ctx context.Context, db *gorm.DB, productType *models.ProductType) error
	DeleteProductType(ctx context.Context, db *gorm.DB, id int) error
}

type productTypeRepository struct {
}

func (r *productTypeRepository) GetProductTypeByName(ctx context.Context, db *gorm.DB, name string) (*models.ProductType, error) {
	db = db.Table(models.ProductTypeTableName).WithContext(ctx)
	var productType models.ProductType
	err := db.Where("name = ?", name).First(&productType).Error
	if err != nil {
		return nil, err
	}
	return &productType, nil
}

func (r *productTypeRepository) CreateProduct(ctx context.Context, db *gorm.DB, productType *models.ProductType) error {
	db = db.Table(models.ProductTypeTableName).WithContext(ctx)
	err := db.Create(productType).Error
	return err
}

func (r *productTypeRepository) GetProductTypeList(ctx context.Context, db *gorm.DB, name string) ([]*models.ProductType, error) {
	db = db.Table(models.ProductTypeTableName).WithContext(ctx)
	var productTypes []*models.ProductType
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	err := db.Find(&productTypes).Error
	if err != nil {
		return nil, err
	}
	return productTypes, nil
}

func (r *productTypeRepository) GetProductTypeById(ctx context.Context, db *gorm.DB, id int) (*models.ProductType, error) {
	db = db.Table(models.ProductTypeTableName).WithContext(ctx)
	var productType models.ProductType
	err := db.Where("id = ?", id).First(&productType).Error
	if err != nil {
		return nil, err
	}
	return &productType, nil
}

func (r *productTypeRepository) UpdateProductType(ctx context.Context, db *gorm.DB, productType *models.ProductType) error {
	db = db.Table(models.ProductTypeTableName).WithContext(ctx)
	err := db.Updates(productType).Error
	return err
}

func (r *productTypeRepository) DeleteProductType(ctx context.Context, db *gorm.DB, id int) error {
	db = db.Table(models.ProductTypeTableName).WithContext(ctx)
	err := db.Where("id = ?", id).Delete(&models.ProductType{}).Error
	return err
}

func NewProductTypeRepository() ProductTypeRepository {
	return &productTypeRepository{}
}

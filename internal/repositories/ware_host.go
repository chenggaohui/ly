package repositories

import (
	"context"
	"gorm.io/gorm"
	"ly/pkg/models"
)

type WareHostRepository interface {
	CreateProduct(ctx context.Context, db *gorm.DB, ware *models.WareHost) error
	GetProductTypeList(ctx context.Context, db *gorm.DB, name string) ([]*models.WareHost, error)
	GetProductTypeById(ctx context.Context, db *gorm.DB, id int) (*models.WareHost, error)
	UpdateProductType(ctx context.Context, db *gorm.DB, ware *models.WareHost) error
	DeleteProductType(ctx context.Context, db *gorm.DB, id int) error
	GetProductTypeByName(ctx context.Context, db *gorm.DB, name string) (*models.WareHost, error)
}

type wareHostRepository struct {
}

func (r *wareHostRepository) GetProductTypeByName(ctx context.Context, db *gorm.DB, name string) (*models.WareHost, error) {
	db = db.Table(models.WareHostTableName).WithContext(ctx)
	var wareHost models.WareHost
	err := db.Where("name = ?", name).First(&wareHost).Error
	if err != nil {
		return nil, err
	}
	return &wareHost, nil
}

func (r *wareHostRepository) CreateProduct(ctx context.Context, db *gorm.DB, ware *models.WareHost) error {
	db = db.Table(models.WareHostTableName).WithContext(ctx)
	err := db.Create(ware).Error
	return err
}

func (r *wareHostRepository) GetProductTypeList(ctx context.Context, db *gorm.DB, name string) ([]*models.WareHost, error) {
	db = db.Table(models.WareHostTableName).WithContext(ctx)
	var wareHosts []*models.WareHost
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	err := db.Find(&wareHosts).Error
	if err != nil {
		return nil, err
	}
	return wareHosts, nil
}

func (r *wareHostRepository) GetProductTypeById(ctx context.Context, db *gorm.DB, id int) (*models.WareHost, error) {
	db = db.Table(models.WareHostTableName).WithContext(ctx)
	var wareHost models.WareHost
	err := db.Where("id = ?", id).First(&wareHost).Error
	if err != nil {
		return nil, err
	}
	return &wareHost, nil
}

func (r *wareHostRepository) UpdateProductType(ctx context.Context, db *gorm.DB, ware *models.WareHost) error {
	db = db.Table(models.WareHostTableName).WithContext(ctx)
	err := db.Updates(ware).Error
	return err
}

func (r *wareHostRepository) DeleteProductType(ctx context.Context, db *gorm.DB, id int) error {
	db = db.Table(models.WareHostTableName).WithContext(ctx)
	err := db.Where("id = ?", id).Delete(&models.WareHost{}).Error
	return err
}

func NewWareHostRepository() WareHostRepository {
	return &wareHostRepository{}
}

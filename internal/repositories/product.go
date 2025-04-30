package repositories

import (
	"context"
	"gorm.io/gorm"
	"ly/pkg/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, db *gorm.DB, product *models.Product) error
	GetProductList(ctx context.Context, db *gorm.DB, product *models.Product, pageNum int, pageSize int) ([]*models.Product, int, error)
	UpdateProduct(ctx context.Context, db *gorm.DB, product *models.Product) error
	GetByProductById(ctx context.Context, db *gorm.DB, productTypeId int) ([]*models.Product, error)
	GetByWareHostId(ctx context.Context, db *gorm.DB, wareHostId int) ([]*models.Product, error)
}

type productRepository struct{}

func (r *productRepository) GetByProductById(ctx context.Context, db *gorm.DB, productTypeId int) ([]*models.Product, error) {
	db = db.Table(models.ProductTableName).WithContext(ctx)
	var products []*models.Product
	err := db.Where("product_type_id = ?", productTypeId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetByWareHostId(ctx context.Context, db *gorm.DB, wareHostId int) ([]*models.Product, error) {
	db = db.Table(models.ProductTableName).WithContext(ctx)
	var products []*models.Product
	err := db.Where("ware_host_id = ?", wareHostId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, db *gorm.DB, product *models.Product) error {
	db = db.Table(models.ProductTableName).WithContext(ctx)
	err := db.Updates(product).Error
	return err
}

func (r *productRepository) GetProductList(ctx context.Context, db *gorm.DB, product *models.Product, pageNum int, pageSize int) ([]*models.Product, int, error) {
	db = db.Table(models.ProductTableName).WithContext(ctx)
	var products []*models.Product
	var total int64
	if product.Name != "" {
		db = db.Where("name = ?", product.Name)
	}
	if product.IsExpired != 0 {
		db = db.Where("is_expired = ?", product.IsExpired)
	}

	if product.ProductTypeId > 0 {
		db = db.Where("product_type_id = ?", product.ProductTypeId)
	}

	if product.WareHostId > 0 {
		db = db.Where("ware_host_id = ?", product.WareHostId)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if pageNum > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	err = db.Order("name desc").Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, int(total), nil
}

func (r *productRepository) CreateProduct(ctx context.Context, db *gorm.DB, product *models.Product) error {
	db = db.Table(models.ProductTableName).WithContext(ctx)
	err := db.Create(product).Error
	return err
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

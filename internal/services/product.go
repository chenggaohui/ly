package services

import (
	"context"
	"gorm.io/gorm"
	"ly/internal/repositories"
	"ly/pkg/models"
	"ly/pkg/vo"
	"strings"
	"time"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *vo.ProductVo) error
	GetProductList(ctx context.Context, product *models.Product, pageNum, pageSize int) ([]*vo.ProductVo, int, error)
	UpdateProduct(ctx context.Context, product *vo.ProductVo) error
}

type productService struct {
	db                *gorm.DB
	productRepository repositories.ProductRepository
}

func (p *productService) UpdateProduct(ctx context.Context, product *vo.ProductVo) error {
	err := p.productRepository.UpdateProduct(ctx, p.db, ConvertProductPo(product))
	return err
}

func (p *productService) GetProductList(ctx context.Context, product *models.Product, pageNum, pageSize int) ([]*vo.ProductVo, int, error) {
	products, total, err := p.productRepository.GetProductList(ctx, p.db, product, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*vo.ProductVo, 0)
	for _, p := range products {
		result = append(result, ConvertProductVo(p))
	}
	return result, total, nil
}

func (p *productService) CreateProduct(ctx context.Context, product *vo.ProductVo) error {
	err := p.productRepository.CreateProduct(ctx, p.db, ConvertProductPo(product))
	return err
}

func (p *productService) checkProduct() {
	productVos, _, err := p.GetProductList(context.Background(), &models.Product{}, 0, 0)
	if err != nil {
		return
	}
	for i := range productVos {
		po := ConvertProductPo(productVos[i])
		productionDate := po.ProductionDate
		lastData := productionDate.Add(24 * time.Hour * time.Duration(po.ShelfLife))
		if lastData.Before(time.Now()) {
			po.IsExpired = -1
			sub := time.Now().Sub(lastData)
			split := strings.Split(sub.String(), ".")
			po.ExpirationDate = split[0] + "s"
			p.UpdateProduct(context.Background(), ConvertProductVo(po))
		} else if po.IsExpired == -1 {
			po.IsExpired = 1
			po.ExpirationDate = " "
			p.UpdateProduct(context.Background(), ConvertProductVo(po))
		}
	}
}

func NewProductService(
	db *gorm.DB,
	productRepository repositories.ProductRepository,
) ProductService {
	impl := &productService{
		db:                db,
		productRepository: productRepository,
	}
	go func() {
		timer := time.NewTicker(time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				impl.checkProduct()
			}
		}
	}()
	return impl
}

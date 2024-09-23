package services

import (
	"ly/pkg/models"
	"ly/pkg/vo"
	"time"
)

func ConvertProductPo(p *vo.ProductVo) *models.Product {
	parse, _ := time.ParseInLocation(time.DateTime, p.ProductionDate, time.Local)
	IsExpired := 1
	if p.IsExpired {
		IsExpired = -1
	}
	return &models.Product{
		Id:             p.Id,
		Name:           p.Name,
		Count:          p.Count,
		ProductionDate: parse,
		ShelfLife:      p.ShelfLife,
		ExpirationDate: p.ExpirationDate,
		IsExpired:      IsExpired,
		Price:          p.Price,
	}
}

func ConvertProductVo(p *models.Product) *vo.ProductVo {
	result := &vo.ProductVo{
		Id:             p.Id,
		Name:           p.Name,
		Count:          p.Count,
		ProductionDate: p.ProductionDate.Format(time.DateTime),
		ShelfLife:      p.ShelfLife,
		ExpirationDate: p.ExpirationDate,
		IsExpired:      p.IsExpired == -1,
		Price:          p.Price,
	}
	return result
}

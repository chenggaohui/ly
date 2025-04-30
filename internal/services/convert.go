package services

import (
	"ly/pkg/models"
	"ly/pkg/vo"
	"time"
)

func ConvertProductPo(p *vo.ProductVo) *models.Product {
	parse, _ := time.ParseInLocation(time.DateTime, p.ProductionDate, time.Local)
	pro := &models.Product{
		Id:             p.Id,
		Name:           p.Name,
		Count:          p.Count,
		ProductionDate: parse,
		ShelfLife:      p.ShelfLife,
		ExpirationDate: p.ExpirationDate,
		IsExpired:      p.IsExpired,
		Price:          p.Price,
		WarnDate:       p.WarnDate,
	}
	if p.ProductType != nil {
		pro.ProductTypeId = p.ProductType.Id
	}
	if p.WareHost != nil {
		pro.WareHostId = p.WareHost.Id
	}
	return pro
}

func ConvertProductVo(p *models.Product) *vo.ProductVo {
	result := &vo.ProductVo{
		Id:             p.Id,
		Name:           p.Name,
		Count:          p.Count,
		ProductionDate: p.ProductionDate.Format(time.DateTime),
		ShelfLife:      p.ShelfLife,
		ExpirationDate: p.ExpirationDate,
		IsExpired:      p.IsExpired,
		Price:          p.Price,
		WarnDate:       p.WarnDate,
	}
	return result
}

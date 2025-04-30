package vo

import "ly/pkg/models"

type ProductVo struct {
	Id             int                 `json:"id"`              // 商品ID
	Name           string              `json:"name"`            // 商品名称
	Count          int                 `json:"count"`           // 数量
	ProductionDate string              `json:"production_date"` // 生产日期
	ShelfLife      int                 `json:"shelf_life"`      // 保质期 (天)
	ExpirationDate string              `json:"expiration_date"` // 过期时间
	IsExpired      int                 `json:"is_expired"`      // 过期标志
	Price          float64             `json:"price"`           // 价格
	WarnDate       int                 `json:"warn_date"`       //提醒时间（天）
	ProductType    *models.ProductType `json:"product_type"`    // 商品类型
	WareHost       *models.WareHost    `json:"ware_host"`       // 仓库名称
}

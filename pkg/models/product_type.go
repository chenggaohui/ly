package models

const ProductTypeTableName = "product_type"

type ProductType struct {
	Id   int    `json:"id"`   // typeID
	Name string `json:"name"` // type名称
}

func (p ProductType) TableName() string {
	return ProductTypeTableName
}

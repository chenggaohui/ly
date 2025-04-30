package models

const WareHostTableName = "ware_host"

type WareHost struct {
	Id   int    `json:"id"`   // wareID
	Name string `json:"name"` // ware名称
}

func (p WareHost) TableName() string {
	return WareHostTableName
}

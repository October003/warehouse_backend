package models

import "gorm.io/gorm"

type ItemSql struct {
	Item
	gorm.Model
}

// Item
type Item struct {
	ItemID          string `json:"item_id" gorm:"unique"` // 物料编号
	ItemName        string `json:"item_name"`             // 物料名称
	Specification   string `json:"specification"`         // 规格型号
	Unit            string `json:"unit"`                  // 计量单位
	StorageLocation string `json:"storage_location"`      // 库位
	OpenNumber      string `json:"open_number"`           // 期初库存
	InNumber        string `json:"in_number"`             // 入库
	OutNumber       string `json:"out_number"`            // 出库
	Number          string `json:"number"`                // 结存
	WarnNumber      string `json:"warn_number"`           // 预警值
}

func (ItemSql) TableName() string {
	return "item_sql"
}

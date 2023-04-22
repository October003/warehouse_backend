package models

import "gorm.io/gorm"

// 入库明细
type InboundSql struct {
	Details
	gorm.Model
}

// 出库明细
type OutboundSql struct {
	Details
	gorm.Model
}

type Details struct {
	ItemID    string `json:"item_id"`   //物料编号
	Person    string `json:"person"`    //申请人
	Number    string `json:"number"`    //数量
	Comment   string `json:"comment"`   //备注
	Timestamp string `json:"timestamp"` //时间戳
}

func (InboundSql) TableName() string {
	return "inbound_sql"
}
func (OutboundSql) TableName() string {
	return "outbound_sql"
}

package models

import "gorm.io/gorm"

// 入库明细
type InboundDetailSql struct {
	Details
	gorm.Model
}

// 出库明细
type OutboundDetailSql struct {
	Details
	gorm.Model
}

type Details struct {
	MaterialID string `json:"item_id"`   //物料编号
	Person     string `json:"person"`    //申请人
	Number     string `json:"number"`    //数量
	Comment    string `json:"comment"`   //备注
	Timestamp  string `json:"timestamp"` //时间戳
}

func (InboundDetailSql) TableName() string {
	return "inbound_detail_sql"
}
func (OutboundDetailSql) TableName() string {
	return "outbound_detail_sql"
}

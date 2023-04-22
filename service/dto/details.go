package dto

type DetailsReq struct {
	ItemID    string `json:"item_id"`   //物料标号
	Person    string `json:"person"`    //申请人
	Number    string `json:"number"`    //数量
	Comment   string `json:"comment"`   //备注
	Timestamp string `json:"timestamp"` //时间戳
}

type DetailQueryReq struct {
	ID             string `json:"id"`
	Timestamp      string `json:"timestamp"`       // 时间戳
	ItemID         string `json:"item_id"`         // 物料编号
	ItemName       string `json:"item_name"`       // 物料名称
	Specification  string `json:"specification"`   // 规格型号
	Unit           string `json:"unit"`            // 计量单位
	StrongLocation string `json:"strong_location"` // 库位
	Number         string `json:"number"`          // 数量
	Person         string `json:"person"`          // 申请人
	Comment        string `json:"comment"`         // 备注
}

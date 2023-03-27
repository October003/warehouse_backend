package dto

type DetailsReq struct {
	MaterialID string `json:"material_id"` //物料标号
	Person     string `json:"person"`      //申请人
	Number     string `json:"number"`      //数量
	Comment    string `json:"comment"`     //备注
	Time       string `json:"time"`        //时间戳
}
type DetailQueryReq struct {
	ID             string `json:"id"`
	Timestamp      string `json:"timestamp"`       // 时间戳
	ItemID         string `json:"material_id"`     // 物料编号
	ItemName       string `json:"material_name"`   // 物料名称
	Specification  string `json:"specification"`   // 规格型号
	Unit           string `json:"unit"`            // 计量单位
	StrongLocation string `json:"strong_location"` // 库位
	Number         string `json:"number"`          // 数量
	Person         string `json:"person"`          // 申请人
	Comment        string `json:"comment"`         // 备注
}

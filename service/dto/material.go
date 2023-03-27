package dto

type Material struct {
	MaterialID      string `json:"material_id" gorm:"unique"` //物料编号
	MaterialName    string `json:"material_name"`             //物料名称
	Specification   string `json:"specification"`             //规格型号
	Unit            string `json:"unit"`                      //计量单位
	StorageLocation string `json:"storage_location"`          //库位
	OpenNumber      string `json:"open_number"`               //期初库存
	InNumber        string `json:"in_number"`                 //入库
	OutNumber       string `json:"out_number"`                //出库
	ResNumber       string `json:"res_number"`                // 库存余结
	WarnNumber      string `json:"warn_number"`               // 预警值
}

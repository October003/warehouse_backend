package models

type Config struct {
	Name string `json:"name" gorm:"name;unique"`
	ModelTime
}

type InBoundPerson Config
type OutBoundPerson Config
type Unit Config
type StorageLocation Config
type MaterialID Config

type AdminConfig struct {
	InBoundPersons  []InBoundPerson   `json:"inbound_persons"`  //入库人
	OutBoundPersons []OutBoundPerson  `json:"outbound_persons"` //出库人
	Units           []Unit            `json:"units"`            //计量单位
	StorageLocation []StorageLocation `json:"strong_locations"` //库位
	MaterialIds     []MaterialID      `json:"item_ids"`         //物料编号
}

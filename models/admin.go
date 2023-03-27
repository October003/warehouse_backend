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
	InBoundPersons  []InBoundPerson   `json:"in_bound_persons"`  //入库人
	OutBoundPersons []OutBoundPerson  `json:"out_bound_persons"` //出库人
	Units           []Unit            `json:"units"`             //计量单位
	StorageLocation []StorageLocation `json:"storage_location"`  //库位
	MaterialIds     []MaterialID      `json:"material_ids"`      //物料编号
}

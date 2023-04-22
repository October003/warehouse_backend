package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Name string `json:"name" gorm:"name;unique"`
}

type InBoundPersons Config
type OutBoundPersons Config
type Units Config
type StrongLocation Config
type ItemIDs Config

type AdminConfig struct {
	InBoundPersons  []InBoundPersons  `json:"inbound_persons"`
	OutBoundPersons []OutBoundPersons `json:"outbound_persons"`
	Units           []Units           `json:"units"`
	StrongLocation  []StrongLocation  `json:"strong_locations"`
	ItemIDs         []ItemIDs         `json:"item_ids"`
}

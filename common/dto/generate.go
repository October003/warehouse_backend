package dto

type ObjectById struct {
	Id  int   `uri:"id"`
	Ids []int `json:"ids"`
}

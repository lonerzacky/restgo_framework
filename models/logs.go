package models

type Logs struct {
	Id         int         `gorm:"primary_key" json:"id,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
	Responses  interface{} `json:"responses,omitempty"`
}

func (Logs) TableName() string {
	return "logs"
}

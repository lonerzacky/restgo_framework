package models

type Sys_role struct {
	Sysrole_kode string `gorm:"primary_key" json:"sysrole_kode,omitempty"`
	Sysrole_nama string `json:"sysrole_nama,omitempty"`
}

func (Sys_role) TableName() string {
	return "sys_role"
}

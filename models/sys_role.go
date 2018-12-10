package models

type SysRole struct {
	SysroleKode string `gorm:"primary_key" json:"sysrole_kode,omitempty"`
	SysroleNama string `json:"sysrole_nama,omitempty"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

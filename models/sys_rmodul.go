package models

type SysRmodul struct {
	SysroleKode  string `gorm:"primary_key" json:"sysrole_kode,omitempty"`
	SysmodulKode string `gorm:"primary_key" json:"sysmodul_kode,omitempty"`
}

func (SysRmodul) TableName() string {
	return "sys_rmodul"
}

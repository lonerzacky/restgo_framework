package models

type SysUser struct {
	SysuserId          string `gorm:"primary_key" json:"sysuser_id,omitempty"`
	SysroleKode        string `json:"sysrole_kode,omitempty"`
	SysuserNama        string `json:"sysuser_nama,omitempty"`
	SysuserPassw       string `json:"sysuser_passw,omitempty"`
	SysuserNamalengkap string `json:"sysuser_namalengkap,omitempty"`
	SysuserEmail       string `json:"sysuser_email,omitempty"`
	SysroleNama        string `sql:"-"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

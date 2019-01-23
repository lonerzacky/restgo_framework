package models

type SysUser struct {
	SysuserId          string `gorm:"primary_key" json:"sysuser_id,omitempty"`
	SysroleKode        string `json:"sysrole_kode,omitempty"`
	SysuserNama        string `json:"sysuser_nama,omitempty"`
	SysuserPassw       string `json:"sysuser_passw,omitempty"`
	SysuserNamalengkap string `json:"sysuser_namalengkap,omitempty"`
	SysuserEmail       string `json:"sysuser_email,omitempty"`
}

type GetSysUser struct {
	SysuserId          string `json:"sysuser_id,omitempty"`
	SysroleKode        string `json:"sysrole_kode,omitempty"`
	SysroleNama        string `json:"sysrole_nama,omitempty"`
	SysuserNama        string `json:"sysuser_nama,omitempty"`
	SysuserPassw       string `json:"sysuser_passw,omitempty"`
	SysuserNamalengkap string `json:"sysuser_namalengkap,omitempty"`
	SysuserEmail       string `json:"sysuser_email,omitempty"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (GetSysUser) TableName() string {
	return "sys_user"
}

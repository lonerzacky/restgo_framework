package models

type Sys_user struct {
	Sysuser_id          string `gorm:"primary_key" json:"sysuser_id,omitempty"`
	Sysrole_kode        string `json:"sysrole_kode,omitempty"`
	Sysuser_nama        string `json:"sysuser_nama,omitempty"`
	Sysuser_passw       string `json:"sysuser_passw,omitempty"`
	Sysuser_namalengkap string `json:"sysuser_namalengkap,omitempty"`
	Sysuser_email       string `json:"sysuser_email,omitempty"`
	Sysrole_nama        string `sql:"-"`
}

func (Sys_user) TableName() string {
	return "sys_user"
}

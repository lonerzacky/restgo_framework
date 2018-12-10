package models

type Sys_modul struct {
	Sysmodul_kode    string `gorm:"primary_key" json:"sysmodul_kode,omitempty"`
	Sysmodul_nama    string `json:"sysmodul_nama,omitempty"`
	Sysmodul_url     string `json:"sysmodul_url,omitempty"`
	Sysmodul_icon    string `json:"sysmodul_icon,omitempty"`
	Sysmodul_parent  string `json:"sysmodul_parent,omitempty"`
	Sysmodul_no_urut string `json:"sysmodul_no_urut,omitempty"`
}

func (Sys_modul) TableName() string {
	return "sys_modul"
}

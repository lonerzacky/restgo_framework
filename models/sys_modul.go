package models

type SysModul struct {
	SysmodulKode   string `gorm:"primary_key" json:"sysmodul_kode,omitempty"`
	SysmodulNama   string `json:"sysmodul_nama,omitempty"`
	SysmodulUrl    string `json:"sysmodul_url,omitempty"`
	SysmodulIcon   string `json:"sysmodul_icon,omitempty"`
	SysmodulParent string `json:"sysmodul_parent,omitempty"`
	SysmodulNoUrut string `json:"sysmodul_no_urut,omitempty"`
}

func (SysModul) TableName() string {
	return "sys_modul"
}

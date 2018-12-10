package controllers

import (
	"../functions"
	"github.com/gin-gonic/gin"
	"net/http"
	"restgo_framework/models"
)

func VerifyLogin(context *gin.Context) {
	var sysUser models.SysUser
	var Username = context.PostForm("Username")
	var Password = context.PostForm("Password")
	var count int
	result := db.
		Select("sysuser_id,sys_role.sysrole_kode, sysrole_nama, sysuser_nama, sysuser_namalengkap, sysuser_email").
		Joins("JOIN sys_role ON sys_role.sysrole_kode = sys_user.sysrole_kode").
		Where("sysuser_nama = ? AND sysuser_passw = ?", Username, Password).
		Find(&sysUser).Count(&count)
	if result.RecordNotFound() {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "LOGIN GAGAL : USERNAME ATAU PASSWORD SALAH", "Record Not Found"))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "LOGIN SUKSES", result.Value))
	}
}

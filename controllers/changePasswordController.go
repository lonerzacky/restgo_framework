package controllers

import (
	"../functions"
	"github.com/gin-gonic/gin"
	"net/http"
	"restgo_framework/models"
)

func ChangePassword(context *gin.Context) {
	var sysUser models.Sys_user
	var SysuserId = context.PostForm("Sysuser_id")
	var OldPassword = context.PostForm("OldPassword")
	var NewPassword = context.PostForm("NewPassword")
	result := db.Select("sysuser_passw").Where("sysuser_id = ?", SysuserId).Find(&sysUser)
	if result.RecordNotFound() {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "NO MATCHING USER ID", "Record Not Found"))
	} else {
		if sysUser.Sysuser_passw == functions.CreateHash(OldPassword) {
			resultUpdate := db.Model(&sysUser).Update("sysuser_passw", functions.CreateHash(NewPassword)).Where("sysuser_id", SysuserId)
			context.JSON(http.StatusOK, functions.GetResponseWithData("00", "UBAH PASSWORD SUKSES", resultUpdate))
		} else {
			context.JSON(http.StatusOK, functions.GetResponseWithData("01", "PASSWORD LAMA TIDAK SAMA", ""))
		}

	}

}

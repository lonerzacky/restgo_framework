package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restgo_framework/functions"
	"restgo_framework/models"
)

func ChangePassword(context *gin.Context) {
	var sysUser models.SysUser
	var SysuserId = context.PostForm("SysuserId")
	var OldPassword = context.PostForm("OldPassword")
	var NewPassword = context.PostForm("NewPassword")
	result := db.Select("sysuser_passw").Where("sysuser_id = ?", SysuserId).Find(&sysUser)
	if result.RecordNotFound() {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "NO MATCHING USER ID", "Record Not Found"))
	} else {
		if sysUser.SysuserPassw == functions.CreateHash(OldPassword) {
			resultUpdate := db.Model(&sysUser).Update("sysuser_passw", functions.CreateHash(NewPassword)).Where("sysuser_id", SysuserId)
			context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "00", "UBAH PASSWORD SUKSES", resultUpdate))
		} else {
			context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "PASSWORD LAMA TIDAK SAMA", ""))
		}

	}

}

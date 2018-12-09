package controllers

import (
	"../functions"
	"github.com/gin-gonic/gin"
	"net/http"
	"restgo_framework/models"
)

func ChangePassword(context *gin.Context) {
	var sys_user models.Sys_user
	var Sysuser_id = context.PostForm("Sysuser_id")
	var OldPassword = context.PostForm("OldPassword")
	var NewPassword = context.PostForm("NewPassword")
	result := db.Select("sysuser_passw").Where("sysuser_id = ?", Sysuser_id).Find(&sys_user)
	if result.RecordNotFound() {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "NO MATCHING USER ID", "Record Not Found"))
	} else {
		if sys_user.Sysuser_passw == functions.CreateHash(OldPassword) {
			resultUpdate := db.Model(&sys_user).Update("sysuser_passw", functions.CreateHash(NewPassword)).Where("sysuser_id", Sysuser_id)
			context.JSON(http.StatusOK, functions.GetResponseWithData("00", "UBAH PASSWORD SUKSES", resultUpdate))
		} else {
			context.JSON(http.StatusOK, functions.GetResponseWithData("01", "PASSWORD LAMA TIDAK SAMA", ""))
		}

	}

}

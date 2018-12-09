package controllers

import (
	"../functions"
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListUser(context *gin.Context) {
	var sys_user []models.Sys_user
	if err := db.Model(&sys_user).Select("sysuser_id,sys_role.sysrole_kode, sysrole_nama, sysuser_nama, sysuser_namalengkap, sysuser_email").Joins("JOIN sys_role ON sys_role.sysrole_kode = sys_user.sysrole_kode").Scan(&sys_user).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "GET LIST USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "GET LIST USER SUKSES", sys_user))
	}
}

func InsertUser(context *gin.Context) {
	var sys_user models.Sys_user
	err := context.Bind(&sys_user)
	sys_user.Sysuser_passw = functions.CreateHash(sys_user.Sysuser_passw)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Create(&sys_user).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "TAMBAH USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "TAMBAH USER SUKSES", sys_user))
	}

}

func UpdateUser(context *gin.Context) {
	var sys_user models.Sys_user
	id := context.Params.ByName("sysuser_id")
	if err := db.Where("sysuser_id = ?", id).First(&sys_user).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE USER GAGAL", err))
		return
	}
	err := context.Bind(&sys_user)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Save(&sys_user).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "UPDATE USER SUKSES", sys_user))
	}

}

func DeleteUser(context *gin.Context) {
	var sys_user models.Sys_user
	id := context.Params.ByName("sysuser_id")
	var result = db.Where("sysuser_id = ?", id).Delete(&sys_user)
	if err := result.Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "DELETE USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "DELETE USER SUKSES", result))
	}

}

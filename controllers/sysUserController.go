package controllers

import (
	"../functions"
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListUser(context *gin.Context) {
	var sysUser []models.Sys_user
	if err := db.Model(&sysUser).Select("sysuser_id,sys_role.sysrole_kode, sysrole_nama, sysuser_nama, sysuser_namalengkap, sysuser_email").Joins("JOIN sys_role ON sys_role.sysrole_kode = sys_user.sysrole_kode").Scan(&sysUser).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "GET LIST USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "GET LIST USER SUKSES", sysUser))
	}
}

func InsertUser(context *gin.Context) {
	var sysUser models.Sys_user
	err := context.Bind(&sysUser)
	sysUser.Sysuser_passw = functions.CreateHash(sysUser.Sysuser_passw)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Create(&sysUser).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "TAMBAH USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "TAMBAH USER SUKSES", sysUser))
	}

}

func UpdateUser(context *gin.Context) {
	var sysUser models.Sys_user
	id := context.Params.ByName("sysuser_id")
	if err := db.Where("sysuser_id = ?", id).First(&sysUser).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE USER GAGAL", err))
		return
	}
	err := context.Bind(&sysUser)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Save(&sysUser).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "UPDATE USER SUKSES", sysUser))
	}

}

func DeleteUser(context *gin.Context) {
	var sysUser models.Sys_user
	id := context.Params.ByName("sysuser_id")
	var result = db.Where("sysuser_id = ?", id).Delete(&sysUser)
	if err := result.Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "DELETE USER GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "DELETE USER SUKSES", result))
	}

}

package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"restgo_framework/database"
	"restgo_framework/functions"
	"restgo_framework/models"
)

var db = database.ConnectDB()

func ListRole(context *gin.Context) {
	var sysRole []models.SysRole
	if err := db.Find(&sysRole).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "GET LIST ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "GET LIST ROLE SUKSES", sysRole))
	}
}

func InsertRole(context *gin.Context) {
	var sysRole models.SysRole
	err := context.Bind(&sysRole)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Create(&sysRole).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "TAMBAH ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "00", "TAMBAH ROLE SUKSES", sysRole))
	}

}

func UpdateRole(context *gin.Context) {
	var sysRole models.SysRole
	id := context.Params.ByName("sysrole_kode")
	if err := db.Where("sysrole_kode = ?", id).First(&sysRole).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "UPDATE ROLE GAGAL", err))
		return
	}
	err := context.Bind(&sysRole)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Save(&sysRole).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "UPDATE ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "00", "UPDATE ROLE SUKSES", sysRole))
	}

}

func DeleteRole(context *gin.Context) {
	var sysRole models.SysRole
	id := context.Params.ByName("sysrole_kode")
	var result = db.Where("sysrole_kode = ?", id).Delete(&sysRole)
	if err := result.Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "DELETE ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "00", "DELETE ROLE SUKSES", result))

	}

}

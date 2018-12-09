package controllers

import (
	"../database"
	"../functions"
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db = database.ConnectDB()

func ListRole(context *gin.Context) {
	var sys_role []models.Sys_role
	if err := db.Find(&sys_role).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "GET LIST ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "GET LIST ROLE SUKSES", sys_role))
	}
}

func InsertRole(context *gin.Context) {
	var sys_role models.Sys_role
	err := context.Bind(&sys_role)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Create(&sys_role).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "TAMBAH ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "TAMBAH ROLE SUKSES", sys_role))
	}

}

func UpdateRole(context *gin.Context) {
	var sys_role models.Sys_role
	id := context.Params.ByName("sysrole_kode")
	if err := db.Where("sysrole_kode = ?", id).First(&sys_role).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE ROLE GAGAL", err))
		return
	}
	err := context.Bind(&sys_role)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Save(&sys_role).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "UPDATE ROLE SUKSES", sys_role))
	}

}

func DeleteRole(context *gin.Context) {
	var sys_role models.Sys_role
	id := context.Params.ByName("sysrole_kode")
	var result = db.Where("sysrole_kode = ?", id).Delete(&sys_role)
	if err := result.Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "DELETE ROLE GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "DELETE ROLE SUKSES", result))

	}

}

package controllers

import (
	"../functions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"restgo_framework/models"
)

func ListModul(context *gin.Context) {
	var sysModul []models.SysModul
	if err := db.Find(&sysModul).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "GET LIST MODUL GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "GET LIST MODUL SUKSES", sysModul))
	}
}

func InsertModul(context *gin.Context) {
	var sysModul models.SysModul
	err := context.Bind(&sysModul)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Create(&sysModul).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "TAMBAH MODUL GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "TAMBAH MODUL SUKSES", sysModul))
	}
}

func UpdateModul(context *gin.Context) {
	var sysModul models.SysModul
	id := context.Params.ByName("sysmodul_kode")
	if err := db.Where("sysmodul_kode = ?", id).First(&sysModul).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE MODUL GAGAL", err))
		return
	}
	err := context.Bind(&sysModul)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Save(&sysModul).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "UPDATE MODUL GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "UPDATE MODUL SUKSES", sysModul))
	}

}

func DeleteModul(context *gin.Context) {
	var sysModul models.SysModul
	id := context.Params.ByName("sysmodul_kode")
	var result = db.Where("sysmodul_kode = ?", id).Delete(&sysModul)
	if err := result.Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithData("01", "DELETE MODUL GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithData("00", "DELETE MODUL SUKSES", result))

	}

}

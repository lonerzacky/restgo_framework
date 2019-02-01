package controllers

import (
	"../functions"
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InsertRmodul(context *gin.Context) {
	var sysRmodul models.SysRmodul
	err := context.Bind(&sysRmodul)
	if err != nil {
		log.Fatal("Error Binding")
	}
	if err := db.Create(&sysRmodul).Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "TAMBAH RMODUL GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "00", "TAMBAH RMODUL SUKSES", sysRmodul))
	}
}

func DeleteRmodul(context *gin.Context) {
	var sysRmodul models.SysRmodul
	kodeRole := context.Params.ByName("kodeRole")
	kodeModul := context.Params.ByName("kodeModul")
	var result = db.Where("sysrole_kode = ? AND sysmodul_kode = ? ", kodeRole, kodeModul).Delete(&sysRmodul)
	if err := result.Error; err != nil {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "01", "DELETE RMODUL GAGAL", err))
	} else {
		context.JSON(http.StatusOK, functions.GetResponseWithDataLogging(context, "00", "DELETE RMODUL SUKSES", result))

	}

}

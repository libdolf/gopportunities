package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func ShowOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func DeleteOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func UpdateOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func ListOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

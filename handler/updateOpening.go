package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "POST opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "GET opening",
	})
}
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "GET opening",
	})
}
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "GET opening",
	})
}
func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "GET opening",
	})
}

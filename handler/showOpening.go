package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrodese/gopportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("opening with id: %s not found", id))
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error finding a opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}

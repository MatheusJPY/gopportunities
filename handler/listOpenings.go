package handler

import (
	"net/http"

	"github.com/MatheusJPY/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendResponseError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendReponseSuccess(ctx, "list-openings", openings)

}

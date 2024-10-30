package handler

import (
	"fmt"
	"net/http"

	"github.com/MatheusJPY/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendResponseError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	//Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendResponseError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	//Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendResponseError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendReponseSuccess(ctx, "delete-opening", opening)
}

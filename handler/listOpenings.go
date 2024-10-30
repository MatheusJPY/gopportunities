package handler

import (
	"net/http"

	"github.com/MatheusJPY/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponseSwagger
// @Failure 500 {object} ErrorResponseSwagger
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendResponseError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendReponseSuccess(ctx, "list-openings", openings)

}

package handler

import (
	"net/http"

	"github.com/MatheusJPY/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendResponseError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendResponseError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
		logger.Infof("opening id:%v role updated", id)
	}
	if request.Company != "" {
		opening.Company = request.Company
		logger.Infof("opening id:%v company updated", id)
	}
	if request.Location != "" {
		opening.Location = request.Location
		logger.Infof("opening id:%v location updated", id)
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
		logger.Infof("opening id:%v remote updated", id)
	}
	if request.Link != "" {
		opening.Link = request.Link
		logger.Infof("opening id:%v link updated", id)
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
		logger.Infof("opening id:%v salary updated", id)
	}

	// Save Opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error update opening: %v", err.Error())
		sendResponseError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendReponseSuccess(ctx, "update-opening", opening)
}

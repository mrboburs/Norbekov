package handler

import (
	"net/http"
	"norbekov/model"

	"github.com/gin-gonic/gin"
)

// @Summary Create Contact
// @Tags Contact
// @Description create contact_post
// @ID create-contact_post
// @Accept  json
// @Produce  json
// @Param input body model.Contact true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/contact/create [post]
//@Security ApiKeyAuth
func (handler *Handler) CreateContactPost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.Contact
	err := ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	Id, err := handler.services.Contact.CreateContactPost(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Data: Id, Message: "DONE"})
}

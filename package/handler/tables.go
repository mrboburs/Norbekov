package handler

import (
	"net/http"
	"norbekov/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create TablePost
// @Tags Table
// @Description create table_post
// @ID create-table_post
// @Accept  json
// @Produce  json
// @Param input body model.TablePost true "table info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/table/create [post]
//@Security ApiKeyAuth
func (handler *Handler) CreateTablePost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.TablePost
	err := ctx.BindJSON(&input)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	Id, err := handler.services.CreateTablePost(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Data: Id, Message: "DONE"})
}

// @Summary Upload Table Image
// @Description Upload Table Image
// @ID upload_img_table
// @Tags   Table
// @Accept  json
// @Produce   json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param        id   path  int     true "Param ID"
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success      200   {object}      ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router   /api/table/upload-img/{id} [PATCH]
//@Security ApiKeyAuth
func (handler *Handler) uploadTableImage(ctx *gin.Context) {
	logrus := handler.logrus
	// homeId := ctx.Param("id")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Errorf("syntax error")
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	imageURL, err := handler.services.UploadImage(file, header, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	effectedRowsNum, err := handler.services.Table.UpdateTableImage(id, imageURL, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Message: "Uploaded", Data: imageURL})
}

// @Summary Update  Table By ID
// @Tags Table
// @Description Update Table by id
// @ID update-table-id
// @Accept  json
// @Produce  json
// @Param        id   path  int     true "Param ID"
// @Param input body model.NewsPost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/table/update/{id} [PUT]
//@Security ApiKeyAuth
func (h *Handler) UpdateTable(ctx *gin.Context) {
	logrus := h.logrus
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Errorf("syntax error")
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var input model.TablePost
	err = ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	effectedRowsNum, err := h.services.Table.UpdateTable(id, input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Message: "Updated", Data: id})

}

// @Summary Get table
// @Tags Table
// @Description get table post
// @ID get-table
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object}  model.TableFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/table/get [GET]
func (h *Handler) GetTableById(ctx *gin.Context) {
	logrus := h.logrus
	id := ctx.Query("id")

	post, err := h.services.Table.GetTableById(id, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"post": post,
	})
}

// @Summary Get Delete table
// @Tags Table
// @Description delete table
// @ID delete-table
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} model.TableFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/table/delete [DELETE]
//@Security ApiKeyAuth
func (h *Handler) DeleteTable(ctx *gin.Context) {

	logrus := h.logrus
	id := ctx.Query("id")

	h.services.Table.DeleteTable(id, logrus)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "deleted",
	})
}

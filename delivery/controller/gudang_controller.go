package controller

import (
	"barang/model"
	"barang/model/dto/response"
	"barang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GudangController struct {
	uc usecase.GudangUsecase
	rg *gin.RouterGroup
}

func (cc *GudangController) registerHandler(ctx *gin.Context) {
	var newGudang model.Gudang
	if err := ctx.ShouldBindJSON(&newGudang); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	err := cc.uc.RegisterGudang(newGudang)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	response.SendSingleResponseCreated(
		ctx,
		newGudang,
		"Success Register new Gudang",
	)
}

func (cc *GudangController) findAllHandler(ctx *gin.Context) {
	gudangs, err := cc.uc.FindAllGudang()
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	var data []any
	data = append(data, gudangs)
	response.SendSingleResponse(
		ctx,
		data,
		"Success Get List Gudang",
	)
}

func (cc *GudangController) updateHandler(ctx *gin.Context) {
	var gudang model.Gudang
	if err := ctx.ShouldBindJSON(&gudang); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	err := cc.uc.UpdateGudang(gudang)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response.SendSingleResponseCreated(
		ctx,
		gudang,
		"Success Update Gudang",
	)
}

func (cc *GudangController) deleteByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.uc.DeleteGudangById(id)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	response.SendSingleResponseCreated(
		ctx,
		http.StatusOK,
		"Success Delete Gudang By Id",
	)
}

func (cc *GudangController) Route() {
	router := cc.rg.Group("/gudang")
	router.POST("", cc.registerHandler)
	router.GET("", cc.findAllHandler)
	router.DELETE("/:id", cc.deleteByIdHandler)
	router.PUT("", cc.updateHandler)
}

func NewGudangController(uc usecase.GudangUsecase, router *gin.Engine) *GudangController {
	return &GudangController{
		uc: uc,
		rg: &router.RouterGroup,
	}
}

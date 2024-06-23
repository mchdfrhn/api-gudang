package controller

import (
	"barang/model/dto/request"
	"barang/model/dto/response"
	"barang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarangController struct {
	uc usecase.BarangUsecase
	rg *gin.RouterGroup
}

func (cc *BarangController) registerHandler(ctx *gin.Context) {
	var newBarang request.BarangRequest
	if err := ctx.ShouldBindJSON(&newBarang); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	err := cc.uc.RegisterBarang(newBarang)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	response.SendSingleResponseCreated(
		ctx,
		newBarang,
		"Success Register new Barang",
	)
}

func (cc *BarangController) findAllHandler(ctx *gin.Context) {
	barangs, err := cc.uc.FindAllBarang()
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	var data []any
	data = append(data, barangs)
	response.SendSingleResponse(
		ctx,
		data,
		"Success Get List Barang",
	)
}

func (cc *BarangController) updateHandler(ctx *gin.Context) {
	var barang request.BarangRequest
	if err := ctx.ShouldBindJSON(&barang); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	err := cc.uc.UpdateBarang(barang)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response.SendSingleResponseCreated(
		ctx,
		barang,
		"Success Update Barang",
	)
}

func (cc *BarangController) deleteByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.uc.DeleteBarangById(id)
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
		"Success Delete Barang By Id",
	)
}

func (cc *BarangController) Route() {
	router := cc.rg.Group("/barang")
	router.POST("", cc.registerHandler)
	router.GET("", cc.findAllHandler)
	router.DELETE("/:id", cc.deleteByIdHandler)
	router.PUT("", cc.updateHandler)
}

func NewBarangController(uc usecase.BarangUsecase, router *gin.Engine) *BarangController {
	return &BarangController{
		uc: uc,
		rg: &router.RouterGroup,
	}
}

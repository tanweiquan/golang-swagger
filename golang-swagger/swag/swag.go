package main

import (
	"net/http"

	_ "swagger/swag/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title l-sample API
// @version 1.0
// @description This is a sample api docs for l-sample.

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

func main() {
	route := gin.Default()
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ctrl := new(Controller)
	route.GET("/api/v1/accounts/:id", ctrl.Get)
	_ = route.Run()
}

type Controller struct{}

// Get
// @Summary Show a account
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID" maximum(10)
// @Param name query string false "Account name" enums("z", "l", "w")
// @Success 200 {object} Account
// @Failure 400,404 {object} HttpError
// @Failure 500 {object} HttpError
// @Failure default {object} HttpError
// @Router /accounts/{id} [get]
func (ctrl Controller) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Account{Name: "zlw"})
}

type Account struct {
	Sex  string `json:"sex" enums:"b,g"`
	Name string `json:"name" example:"zlw"`
}

type HttpError struct {
	Code int    `json:"code" example:"1001"`
	Msg  string `json:"msg" example:"server error"`
}

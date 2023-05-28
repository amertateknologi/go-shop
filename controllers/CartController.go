package controllers

import (
	"net/http"

	"github.com/amerta-teknologi/go-shop/models"
	"github.com/amerta-teknologi/go-shop/utils"

	"github.com/gin-gonic/gin"
)

type CartController struct{}

func (ctrl *CartController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.Cart{}
		resp := model.FindAll(c)

		c.HTML(http.StatusOK, "pages/carts", gin.H{
			"webview": utils.Webview,
			"menus":   utils.Menus,
			"carts":   resp.Data,
		})
	}
}

func (ctrl *CartController) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.Cart{}
		model.Post(c)

		c.Redirect(http.StatusFound, "/carts")
	}
}
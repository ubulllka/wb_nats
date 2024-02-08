package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ubulllka/wbL0/internal/cache"
	"github.com/ubulllka/wbL0/internal/pub"
)

func GetAll(ctx *gin.Context) {
	var arr []interface{}
	cache.CACHE.Range(func(k, v interface{}) bool {
		arr = append(arr, k)
		return true
	})
	ctx.JSON(http.StatusOK, &arr)
}

func Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if order, err := cache.GetOrderFromCache(id); err != nil {
		ctx.JSON(http.StatusNotFound, &gin.H{"message": "data not found"})
	} else {
		ctx.JSON(http.StatusOK, &order)
	}
}

func AddFromPub(ctx *gin.Context) {
	id, err := pub.Connect()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError , err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "id: " + id})
}

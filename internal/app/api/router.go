package api

import (
	"github.com/gin-gonic/gin"
	"go-rest-mongo/internal/app/api/controller"
)

func Init() *gin.Engine {
	// Get a user resource
	router := gin.Default()
	router.GET("/txs", controller.TxsList)
	router.GET("/tx/:txhash", controller.TxHash)
	router.GET("/nfts", controller.NftList)
	//router.GET("/users/:id", uc.GetUser)
	//router.DELETE("/users/:id", uc.RemoveUser)
	//router.POST("/users", uc.CreateUser)
	//router.PUT("/users/:id", uc.UpdateUser)
	return router
}

package controller

import (
	"github.com/gin-gonic/gin"
	"go-restful/internal/model"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//findall
func NftList(c *gin.Context) {
	var results []model.Nftt
	err := db.C("nft").Find(bson.M{}).All(&results)
	if err != nil {
		log.Fatal("txs doesn't exist ")
		return
	}

	//c.JSON(200, results)
	c.JSON(200, gin.H{
		"txs": results,
	})

}

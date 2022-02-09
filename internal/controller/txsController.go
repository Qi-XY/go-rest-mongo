package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-mongo/config"
	"go-rest-mongo/internal/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const DB_COLLECTION = "txs"

// Get DB from Mongo Config
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

var db mgo.Database = *MongoConfig()

//findtxs

func TxsList(c *gin.Context) {
	//db := *MongoConfig()
	//fmt.Println("MONGO RUNNING: ", db)

	var results []model.Txss
	err := db.C(DB_COLLECTION).Find(bson.M{}).All(&results)
	if err != nil {
		log.Fatal("txs doesn't exist ")
		return
	}
	c.JSON(200, gin.H{
		"txs": &results,
	})
}

//findall
func TxHash(c *gin.Context) {

	txhash := c.Param("txhash") // Get Param
	fmt.Println(txhash)

	result := model.Txss{}
	err := db.C(DB_COLLECTION).Find(bson.M{"txhash": txhash}).One(&result)
	if err != nil {
		//log.Fatal("txhash doesn't exist ")
		//return
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"txhash": &result.Txhash,
		"height": &result.Height,
		"tx":     &result.Tx,
	})

}

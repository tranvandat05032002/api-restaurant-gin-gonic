package main

import (
	"context"
	"fmt"
	"gin-gonic-gom/Controllers"
	"gin-gonic-gom/Services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	server      *gin.Engine
	us          Services.AccountService
	uc          Controllers.AccountController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

const mongoCnn = "mongodb+srv://tranvandatevondev0503:35701537scss@cluster0.8dgoa.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(mongoCnn)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("MeteorDB").Collection("accounts")
	us = Services.NewUserService(userc, ctx)
	uc = Controllers.New(us)
	server = gin.Default()
}
func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterAccountRoutes(basepath)

	log.Fatal(server.Run(":8080"))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

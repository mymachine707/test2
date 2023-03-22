package main

import (
	"context"
	"fmt"
	"log"

	"exemple.com/sarang-apis/controllers"
	"exemple.com/sarang-apis/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	monogoclient   *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:2717")

	monogoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}

	err = monogoclient.Ping(ctx, readpref.Primary())

	fmt.Println("mongo connection established")

	usercollection = monogoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()

}

func main() {
	defer monogoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	usercontroller.RegistrUserRoutses(basepath)
	log.Fatal(server.Run(":9090"))
}

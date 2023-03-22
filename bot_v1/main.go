package main

import (
	"context"
	"fmt"
	"log"

	"ucode/telegram-bot/controllers"
	"ucode/telegram-bot/services"

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

	botToken := "5948366832:AAEfwTE1B9ji1X-4_w1RXMITBqqt22gK06Y"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken

	offset := 0
	c := 0
	for {
		updates, err := services.GetUpdates(botUrl, offset)
		if err != nil {
			log.Println("smt went wrong: ", err.Error())
		}
		fmt.Println("Kotta for: ", c, updates)
		c++
		l := 0

		for _, update := range updates {
			_ = services.Respond(botUrl, update)
			fmt.Println("ichki for: ", l, update)
			l++
			err = services.Responds(usercollection, update)
			if err != nil {
				log.Println("smt went wrong: ", err.Error())
			}
			offset = update.UpdateId + 1
		}
	}
}

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"ucode/telegram-bot/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetUpdates(botUrl string, offset int) ([]models.Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restRespnse models.RestRespnse
	err = json.Unmarshal(body, &restRespnse)
	if err != nil {
		return nil, err
	}
	return restRespnse.Result, nil
}

func Respond(botUrl string, update models.Update) error {
	var botMessage models.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func Responds(collection *mongo.Collection, update models.Update) error {
	fmt.Println(update)
	_, err := collection.InsertOne(context.Background(), update)
	if err != nil {
		fmt.Println("check")
		return err
	}
	return nil
}

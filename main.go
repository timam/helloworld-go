package main

import (
	"encoding/json"
	"fmt"
	"github.com/timam/helloworld-go/logger"
	"net/http"
)

type SomeStruct struct {
	Color	string	`json:"color"`
	Message string 	`json:"message"`
}

var log = "| http-nio-8080-exec-6 |  INFO | c.b.a.common.aspect.AccessLoggingAspect  |  2029311C58041113 | REST request on com.bkash.autobot.gateway.controller.IntentController.create() with [IntentDto.CreateIntentDto(super=IntentDto(invoice=31f143a76866402981, type=SUBSCRIPTION, subscriptionData=SubscriptionData(amount=15.0, cycle=WEEKLY, serviceReference=MM_CONTEST_15TK_1W, merchantReference=MoMAGIC)))]"

func colorAndMessage(w http.ResponseWriter, r *http.Request) {
	data := SomeStruct{
		Color: "#fc5c65", //FUSION RED
		Message: "kubernetes is awesome",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	logger.Log.Printf(log)
	fmt.Println(log)
}

func main() {
	http.HandleFunc("/", colorAndMessage)
	http.ListenAndServe(":8080", nil)
}

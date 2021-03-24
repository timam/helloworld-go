package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var method = "GET"
var endpoint = "https://eventstream.uatcapp.bka.sh/event-stream/actuator/health"

func gimmeResponse(method, endpoint string) (int, string) {

	request, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		fmt.Println("Request Error :", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		print("Response Error", err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	return response.StatusCode, string(body)
}

func putStatusCodeToCloudwatch(statusCode int)  {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new cloudwatch client.
	svc := cloudwatch.New(sess)

	_ := svc

}

func main() {
	statusCode, body := gimmeResponse(method, endpoint)
	fmt.Println(statusCode, body)
}

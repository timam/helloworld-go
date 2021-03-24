package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"io/ioutil"
	"net/http"
)

var method = "GET"
var endpoint = "https://eventstream.uatcapp.bka.sh/event-stream/actuator/health"
var serviceName = "1test"

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

func putStatusCodeToCloudwatch(status float64, serviceName string) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new cloudwatch client.
	svc := cloudwatch.New(sess)

	_, err := svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("EKS-UPTIME"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("Status"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(status),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("ServiceName"),
						Value: aws.String(string(serviceName)),
					},
					&cloudwatch.Dimension{
						Name:  aws.String("ServiceName2"),
						Value: aws.String(string(serviceName)),
					},
				},
			},
		},
	})

	if err != nil{
		fmt.Println("Cloudwatch Put Error :", err)
	}

}



func main() {
	statusCode, body := gimmeResponse(method, endpoint)

	var statusUp float64
	statusUp = 1.0
	putStatusCodeToCloudwatch(statusUp, serviceName)

	fmt.Println(statusCode, body)
}

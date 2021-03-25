package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)


const statusUp  = 1.0
const statusDown = 0.0

//https://admin.uatsm.bka.sh/admin/actuator/health
//https://callback.uatsm.bka.sh/cps/actuator/health
//https://callback.uatsm.bka.sh/processor/actuator/health/
//https://admin-app.uatsm.bka.sh/admin-app/actuator/health

var method = "GET"
var endpoint = "https://eventstream.uatcapp.bka.sh/event-stream/actuator/health/"
var serviceName = "mysteram"


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

func putStatusCodeToCloudwatch(status float64, serviceName string, httpStatus string) {

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
						Value: aws.String(serviceName),
					},
					&cloudwatch.Dimension{
						Name:  aws.String("StatusCode"),
						Value: aws.String(httpStatus),
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

	for {
		statusCode, body := gimmeResponse(method, endpoint)
		strStatusCode := strconv.Itoa(statusCode)

		if statusCode == 200{
			putStatusCodeToCloudwatch(statusUp, serviceName, strStatusCode)
			log.Println(fmt.Sprintf("%s is: up, status code:  %s, body: %s", serviceName, strStatusCode, body))
		} else {
			putStatusCodeToCloudwatch(statusDown, serviceName, strStatusCode)
			log.Println(fmt.Sprintf("%s is: unhealthy, status code:  %s", serviceName, strStatusCode))
		}
		time.Sleep(1000 *time.Millisecond)
	}

}

package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const ServiceUrl = "https://api.us-west-1.saucelabs.com/v1/magnificent/"

//interval between  consecutive requests in seconds
const requestInterval = 2

//service health monitoring interval in minutes
const serviceHealthInterval = 10

//log file for monitoring service
const logFileName = "service_monitor.log"

func main() {
	log.Print("Starting service monitor")

	f, err := os.OpenFile(logFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	//requestCount is incremented on ever request while successCount is incremented only on succesful request.
	//these parameters are then used to determine health of service
	requestCount, successCount := 0, 0

	// determine service health for serviceHealthinterval minutes
	ticker := time.NewTicker(serviceHealthInterval * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				successRate := successCount * 100 / requestCount
				log.Println("Last ", serviceHealthInterval, " minutes:", successRate, "% of requests successful")
				requestCount, successCount = 0, 0
				if successRate < 25 {
					log.Println("service unavailable")
				}
				log.Println()
			}
		}
	}()

	// make requests to server and write logs for each request
	for {
		response, err := http.Get(ServiceUrl)
		if err != nil {
			log.Println("failed to connect to service", err)
			return
		}
		requestCount++
		if response.StatusCode != http.StatusOK {
			log.Println("service returned error code:", response.StatusCode)
		} else {
			log.Println("service available")
			successCount++
		}
		time.Sleep(requestInterval * time.Second)
	}
}

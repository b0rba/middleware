package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/b0rba/middleware/my-middleware/client/distribution/proxies"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// runExperiment run the experiment.
func runExperiment(waitGroup *sync.WaitGroup, calcValues *utils.CalcValues, numberOfCalls int, start int) {
	defer waitGroup.Done()
	// getting the clientproxy
	namingServer := proxies.InitServer("localhost")
	echo := namingServer.Lookup("Ech").(proxies.EchoProxy)

	// executing
	for i := 0; i < numberOfCalls; i++ {
		initialTime := time.Now() //calculating time
		result := echo.Ech(i + start)
		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		fmt.Println(result)                                            // making the request
		utils.AddValue(calcValues, endTime)                            // pushing to the stored values
		time.Sleep(10 * time.Millisecond)                              // setting the sleep time
	}
}

// goSleep just sleep so the client can mak the requests
func goSleep() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		i--
	}
}

func main() {
	numberOfCalls := 10000
	perCall := 500
	aux := numberOfCalls / perCall
	// creating the calcvalues object
	calcValues := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls))
	var waitGroup sync.WaitGroup
	go goSleep()

	for i := 0; i < aux; i++ {
		waitGroup.Add(1)
		go runExperiment(&waitGroup, &calcValues, numberOfCalls, i * perCall)
		waitGroup.Wait()
	}

	// evaluating
	average := utils.CalcAverage(&calcValues)
	stdv := utils.CalcStandardDeviation(&calcValues, average)

	utils.PrintEvaluation(average, stdv, 8)

}

package main

import (
	"fmt"

	"net/rpc"
	"strconv"
	"time"

	"github.com/b0rba/middleware/utils"
)

func main() {
	numberOfCalls := 10000 // the number of server calls

	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls)) // object to store the rtts

	var reply int
	// connect to servidor
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "O Servidor não está pronto")

	// make requests
	for i := 0; i < numberOfCalls; i++ {
		// prepara request
		args := fmt.Sprintf("%s:%d", "string number:", i)

		initialTime := time.Now()
		// envia request e recebe resposta
		client.Call("Echo.Echo", args, &reply)

		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		utils.AddValue(&calc, endTime)

		fmt.Printf("%v\n", *&reply)

		time.Sleep(25 * time.Millisecond)
	}
	// evaluating
	avrg := utils.CalcAverage(&calc)
	stdv := utils.CalcStandardDeviation(&calc, avrg)

	utils.PrintEvaluation(avrg, stdv, 8)
}

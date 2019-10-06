package utils

import (
	"fmt"
	"log"
	"math"
)

// PrintError is a function to print an error message.
//
// Parameters:
//  err     - the error.
//	message - the string with extra information.
//
// Returns:
//	none
//
func PrintError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

// CalcValues is a structure to hold time information for calculation.
//
// Members:
//  Used   - number of used values.
//  Values - list with stored values.
//
type CalcValues struct {
	Used   int
	Values []float64
}

// CalcStandardDeviation is a function to calculate the standard deviation of a CalcValues object.
//
// Parameters:
//  calc    - pointer to the CalcValues.
//  average - average value of CalcValues values.
//
// Returns:
//  the standard deviation.
//
func CalcStandardDeviation(calc *CalcValues, average float64) float64 {
	var sd float64
	for i := 0; i < calc.Used; i++ {
		sd += math.Pow(calc.Values[i]-average, 2)
	}
	sd = math.Sqrt(sd / float64(calc.Used))
	return sd
}

// CalcAverage is a function to calculate the average value of a CalcValues object.
//
// Parameters:
//  calc - pointer to the CalcValues.
//
// Returns:
//  the average value.
//
func CalcAverage(calc *CalcValues) float64 {
	var total float64
	for i := 0; i < calc.Used; i++ {
		total += calc.Values[i]
	}
	return total / float64(calc.Used)
}

// AddValue is a function to add a value to a CalcValues object.
//
// Parameters:
//  calc  - pointer to the CalcValues.
//  value - the new value.
//
// Returns:
//  none
//
func AddValue(calc *CalcValues, value float64) {
	calc.Values[calc.Used] = value
	calc.Used++
}

// InitCalcValues is a function to initialize a CalcValues object.
//
// Parameters:
//  values - list of values.
//
// Returns:
//  a CalcValues object.
//
func InitCalcValues(values []float64) CalcValues {
	calc := CalcValues{Used: 0, Values: values}
	return calc
}

// PrintEvaluation is a function to print the average value and standard deviation.
//
// Parameters:
//  avrg - the average value.
//  stdv - the standard deviation.
//
// Returns:
//  none
//
func PrintEvaluation(avrg float64, stdv float64) {
	fmt.Printf("The Average RTT was: %0.4fms.\n", avrg)
	fmt.Printf("The Standard Deviation on the RTT was: %0.4fms.\n", stdv)
}

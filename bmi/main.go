package main

import (
	"github.com/yabuking84/golang-experiments/bmi/calculate"
	"github.com/yabuking84/golang-experiments/bmi/info"
	"github.com/yabuking84/golang-experiments/bmi/metrics"
)

func main() {

	info.PrintInfo()

	weight, height := metrics.GetMetrics()

	info.PrintMetrics(weight, height)

	bmi := calculate.GetBMI(weight, height)

	info.PrintBMI(bmi)

}

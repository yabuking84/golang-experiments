package metrics

import (
	"fmt"

	"github.com/yabuking84/golang-experiments/bmi/info"
	"github.com/yabuking84/golang-experiments/bmi/readerhandler"
)

func GetMetrics() (float64, float64) {

	fmt.Print(info.PromptWeight)
	weight := readerhandler.GetFloatInput()
	fmt.Print(info.PromptHeight)
	height := readerhandler.GetFloatInput()

	return weight, height

}

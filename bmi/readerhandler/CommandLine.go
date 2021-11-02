package readerhandler

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetFloatInput() float64 {
	var retVal float64
	input, _ := reader.ReadString('\n')
	input = cleanString(input)
	retVal = toFloat(input)
	return retVal
}

func cleanString(s string) string {
	var retVal string
	if runtime.GOOS == "windows" {
		retVal = strings.Replace(s, "\r\n", "", -1)
	} else {
		retVal = strings.Replace(s, "\n", "", -1)
	}
	return retVal
}

func toFloat(s string) float64 {
	var retVal float64

	retVal, _ = strconv.ParseFloat(s, 64)

	// asd := fmt.Sprintf("%f", retVal)
	// fmt.Println("this is " + asd)

	return retVal
}

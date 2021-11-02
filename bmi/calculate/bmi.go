package calculate

import "math"

func GetBMI(weight float64, height float64) float64 {
	retVal := weight / (height * height)
	retVal = math.Round(retVal*10000) / 10000

	return retVal
}

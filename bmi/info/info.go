package info

import "fmt"

const MainTitle = "BMI  Calculator"
const Separator = "-----------------------------------"
const PromptWeight = "Enter weight (kg): "
const PromptHeight = "Enter height (m): "

func PrintInfo() {
	fmt.Println(MainTitle)
	fmt.Println(Separator)
}

func PrintMetrics(weight float64, height float64) {
	fmt.Printf("Your weight is %vkg and your height is %vm.", weight, height)
	fmt.Println()
}

func PrintBMI(bmi float64) {
	squared := '\u00B2'
	fmt.Printf("Your BMI is %v kg/m%c.", bmi, squared)
	fmt.Println()
}

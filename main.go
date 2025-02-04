package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	println("__IMT Calculator __")
	for {
		UserWeight, UserHeight := getUserInput()
		IMT, err := calculateIMT(UserWeight, UserHeight)
		if err != nil {
			panic("Calculation parameters are not specified")
		}
		outputResult(IMT)
		RepeatCalculation := checkCalculationAgain()
		if !RepeatCalculation {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Your Body Mass Index: %.0f", IMT)
	fmt.Println(result)
	switch {
	case IMT <= 16:
		fmt.Println("You are severely underweight")
	case IMT < 18.5:
		fmt.Println("You are underweight")
	case IMT < 25:
		fmt.Println("You are a normal weight")
	case IMT < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("You have a degree of obesity")
	}
}

func calculateIMT(UserWeight float64, UserHeight float64) (float64, error) {
	if UserWeight <= 0 || UserHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := UserWeight / math.Pow(UserHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var UserWeight, UserHeight float64
	fmt.Printf(` Enter your height (cm):
	`)
	fmt.Scan(&UserHeight)
	fmt.Print(` Enter your weight:
	`)
	fmt.Scan(&UserWeight)
	return UserWeight, UserHeight
}

func checkCalculationAgain() bool {
	fmt.Println("Do you want to check IMT again? (y/n)")
	var Option string
	fmt.Scan(&Option)
	if Option == "y" || Option == "Y" {
		return true
	}
	return false
}

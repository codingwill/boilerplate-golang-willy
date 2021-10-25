package services

import (
	"math/rand"
	"projectstructuring/models"
)

func ServiceGetRandomizeNumber(numbers models.RandomArrayInteger, option string) (result models.Response) {
	result.Status = 200
	result.From = "system"
	result.Message = "Fetching data success"
	if option == "method" {
		result.Data = numbers.Randomize()
	} else if option == "function" {
		// Fisher-Yates algo
		for i := len(numbers.Data) - 1; i >= 1; i-- {
			j := rand.Intn(i + 1)
			numbers.Data[i], numbers.Data[j] = numbers.Data[j], numbers.Data[i]
		}
		result.Data = numbers.Data
	}
	return result
}

func ServiceGetRandomizeString(words *models.RandomArrayString, option string) (result models.Response) {
	result.Status = 200
	result.From = "system"
	result.Message = "Fetching data success"
	if option == "method" {
		result.Data = words.Randomize()
	} else if option == "function" {
		// Fisher-Yates algo
		for i := len(words.Data) - 1; i >= 1; i-- {
			j := rand.Intn(i + 1)
			words.Data[i], words.Data[j] = words.Data[j], words.Data[i]
		}
		result.Data = words.Data
	}

	return result
}

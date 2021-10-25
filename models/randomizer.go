package models

import "math/rand"

type RandomArrayInteger struct {
	Data []int32 `json:"data"`
}

type RandomArrayString struct {
	Data []string `json:"data"`
}

func (number RandomArrayInteger) Randomize() (result []int32) {
	// Fisher-Yates algo
	for i := len(number.Data) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		number.Data[i], number.Data[j] = number.Data[j], number.Data[i]
	}
	result = number.Data
	return result
}

func (words RandomArrayString) Randomize() (result []string) {
	// Fisher-Yates algo
	for i := len(words.Data) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		words.Data[i], words.Data[j] = words.Data[j], words.Data[i]
	}
	result = words.Data
	return result
}

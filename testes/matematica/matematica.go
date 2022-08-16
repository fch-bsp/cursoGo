package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcuilar média


func Media(numeros...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2F", media), 64)
	return mediaArredondada
}
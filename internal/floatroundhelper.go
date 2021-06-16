package internal

import "strconv"
import "fmt"

func RoundFloat(x float32) float32 {
	i := fmt.Sprintf("%.2f", x)
	f, _ := strconv.ParseFloat(i, 2)
	return float32(f)
}
package chap06

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSeveral(t *testing.T) {
	max := findMax(1.2, 2.3, 3.4, 4.5)
	assert.Equalf(t, max, 4.5, "error message %s", "formatted")
}

func TestSlice(t *testing.T) {
	float64s := []float64{1.2, 2.3, 3.4, 4.5}
	max := findMax(float64s...)
	assert.Equal(t, max, 4.5)
}

func findMax(values ...float64) float64 {
	max := math.Inf(-1)
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

package chap11

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type OverHeatError float64

func (overHeatError OverHeatError) Error() string {
	return fmt.Sprintf("overheating tmp, value is %f", overHeatError)
}

func checkTemperature(actual float64, safe float64) error {
	if actual > safe {
		return OverHeatError(actual - safe)
	}
	return nil
}

func TestCustomError(t *testing.T) {
	err := checkTemperature(1.2, 2.3)
	assert.Nil(t, err)

	err = checkTemperature(2.3, 1.2)
	assert.NotNil(t, err)
	t.Log(err)
}

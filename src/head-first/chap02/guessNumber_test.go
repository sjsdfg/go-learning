package chap02

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGuessNumber(t *testing.T) {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100)
	guessNumber := rand.Intn(60)

	for target != guessNumber {
		guessNumber += isLower(t, target, guessNumber)
	}

	assert.Equal(t, target, guessNumber)
}

func isLower(t *testing.T, target int, actual int) int {
	t.Logf("target is %d, actual is %d", target, actual)
	if target > actual {
		t.Log("target bigger than actual")
		return 1
	} else if target < actual {
		t.Log("target lower than actual")
		return -1
	}
	t.Log("target equals actual")
	return 0
}

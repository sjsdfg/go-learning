package chap02

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomWithNoSeed(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(rand.Intn(20))
	}
}

func TestRandomWithSeed(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		t.Log(rand.Intn(20))
	}
}

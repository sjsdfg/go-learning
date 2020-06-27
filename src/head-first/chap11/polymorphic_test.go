package chap11

import (
	"fmt"
	"testing"
)

type NoiseMaker interface {
	MakeNoise()
}

type Whistle string

func (whistle Whistle) MakeNoise() {
	fmt.Println("Whistle MakeNoise", whistle)
}

type Robot string

func (robot Robot) MakeNoise() {
	fmt.Println("Robot MakeNoise", robot)
}

func TestPoly(t *testing.T) {
	var noiseMaker NoiseMaker
	noiseMaker = Whistle("123")
	noiseMaker.MakeNoise()

	noiseMaker = Robot("robot")
	noiseMaker.MakeNoise()
}

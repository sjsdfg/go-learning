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

func (robot Robot) Walk() {
	fmt.Println("Robot is walking", robot)
}

func TestPoly(t *testing.T) {
	var noiseMaker NoiseMaker
	noiseMaker = Whistle("123")
	noiseMaker.MakeNoise()

	noiseMaker = Robot("robot")
	noiseMaker.MakeNoise()
}

func TestTypeAssert(t *testing.T) {
	var noiseMaker NoiseMaker
	noiseMaker = Robot("robot")
	noiseMaker.MakeNoise()

	var robot = noiseMaker.(Robot)
	robot.Walk()
}

func TestTypeAssertFailure(t *testing.T) {
	var noiseMaker NoiseMaker
	noiseMaker = Whistle("123")
	noiseMaker.MakeNoise()

	var robot, ok = noiseMaker.(Robot)
	if ok {
		robot.Walk()
	} else {
		t.Log(ok)
	}
}

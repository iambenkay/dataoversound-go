package main

import (
	aud "github.com/iambenkay/dataoversound-go/audio-utils"
	"math"
	"time"
)

const sampleRate = 44100

func startBroadcaster(){
	for {
		makeSound(1000, sampleRate)
		time.Sleep(300 * time.Millisecond)
	}
}

func makeSound(Feed, sampleRate float64) (ac *AudioContext) {
	ac = &AudioContext{nil, Feed / sampleRate, 0, Feed / sampleRate, 0}

	var err error

	ac.Stream, err = aud.OpenDefaultStream(0, 2, sampleRate, 0, ac.process)
	aud.PanicIfErr(err)
	defer aud.PanicIfErrDeferred(ac.Close)

	aud.PanicIfErr(ac.Start())
	time.Sleep(10000 * time.Millisecond)
	aud.PanicIfErr(ac.Stop())
	return
}

type AudioContext struct {
	*aud.Stream
	stepL, phaseL, stepR, phaseR float64
}

func (ac *AudioContext) process(out [][]float32) {
	// TODO come up with a better way to smoothen out the figures with realistic values.
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * ac.phaseL))
		_, ac.phaseL = math.Modf(ac.phaseL + ac.stepL)
		out[1][i] = float32(math.Sin(2 * math.Pi * ac.phaseR))
		_, ac.phaseR = math.Modf(ac.phaseR + ac.stepR)
	}
}

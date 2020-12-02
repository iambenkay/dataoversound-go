package main

import (
	"fmt"
	aud "github.com/iambenkay/dataoversound-go/audio-utils"
	"os"
	"os/signal"
	"time"
)

const sampleRateRecorder = 44100

var (
	interrupt chan os.Signal
)

func startSubscriber() {
	fmt.Println("Recording.  Press Ctrl-C to stop.")

	interrupt = make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	listen()
}

func listen() {
	in := make([]int32, 64)
	stream, err := aud.OpenDefaultStream(1, 0, sampleRateRecorder, len(in), in)
	aud.PanicIfErr(err)

	defer aud.PanicIfErrDeferred(stream.Close)

	aud.PanicIfErr(stream.Start())

	go func() {
		for {
			time.Sleep(2 * time.Second)
			// TODO Actually understand how to interpret data received
			fmt.Println(avg(in...))
		}
	}()

	for {
		aud.PanicIfErr(stream.Read())
		select {
		case <-interrupt:
			aud.PanicIfErr(stream.Stop())
			return
		default:
		}
	}
}

func avg(nums ...int32) int32 {
	var total int32 = 0
	var i = 0
	var n int32 = 0
	for i, n = range nums {
		total += n
	}
	return total/int32(i + 1)
}

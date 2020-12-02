package main

import (
	aud "github.com/iambenkay/dataoversound-go/audio-utils"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: ./dataoversound-go <mode> NB: mode is one of broadcaster or subscriber")
		os.Exit(1)
	}
	mode := os.Args[1]
	if len(mode) == 0 || !(strings.EqualFold(mode, "broadcaster") || strings.EqualFold(mode, "subscriber")) {
		println("Usage: ./dataoversound-go <mode> NB: mode is one of broadcaster or subscriber")
		os.Exit(1)
	}
	aud.PanicIfErr(aud.Initialize())
	defer aud.PanicIfErrDeferred(aud.Terminate)

	switch mode {
	case "broadcaster":
		startBroadcaster()
		break
	case "subscriber":
		startSubscriber()
	}
}

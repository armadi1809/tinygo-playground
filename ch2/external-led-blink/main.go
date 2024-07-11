package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	yellowLED := machine.D13
	yellowLED.Configure(outputConfig)

	for {
		yellowLED.Low()
		time.Sleep(250 * time.Millisecond)
		yellowLED.High()
		time.Sleep(500 * time.Millisecond)

	}

}

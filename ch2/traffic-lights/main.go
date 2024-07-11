package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	redLed := machine.D13
	redLed.Configure(outputConfig)

	yellowLed := machine.D12
	yellowLed.Configure(outputConfig)

	greenLed := machine.D11
	greenLed.Configure(outputConfig)

	for {
		redLed.High()
		time.Sleep(time.Second)
		yellowLed.High()
		time.Sleep(time.Second)

		redLed.Low()
		yellowLed.Low()
		greenLed.High()
		time.Sleep(time.Second)

		greenLed.Low()
		yellowLed.High()
		time.Sleep(time.Second)
		yellowLed.Low()

	}
}

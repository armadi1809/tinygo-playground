package main

import (
	"machine"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	inputConfig := machine.PinConfig{Mode: machine.PinInput}

	led := machine.D13
	led.Configure(outputConfig)

	buttonInput := machine.D2
	buttonInput.Configure(inputConfig)

	for {
		if buttonInput.Get() {
			led.High()
			continue
		}
		led.Low()
	}

}

package main

import (
	"machine"
	"strconv"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	flashCount := 0

	// using analog input(-capable) pin as a digital input because to colocate all pins on the same
	// side of Arduino Uno so I can use just a single screw shield and thus make it more physically smol.
	flashDetector := machine.ADC0
	flashDetector.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	flashDetector.SetInterrupt(machine.PinToggle, func(p machine.Pin) {
		led.Set(true)
		if p.Get() { // don't count both low and high transitions
			flashCount++
		}
	})

	for {
		// string prefix before the number so it's easier to gauge on the receiving end if the line
		// came through intact
		println("num=" + strconv.Itoa(flashCount))
		time.Sleep(2 * time.Second)
		led.Set(false)
	}
}

package main

import (
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/adaptors"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor(adaptors.WithGpiodAccess())
	led_red := gpio.NewLedDriver(r, "15")
	led_grn := gpio.NewLedDriver(r, "36")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led_red.Toggle()
			led_grn.On()
			if led_red.State() {
				led_grn.Off()
			}
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led_red, led_grn},
		work,
	)

	robot.Start()
}

package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	myBuffer := make([]byte, 9999)
	work := func() {
		drone.TakeOff()

		gobot.After(5*time.Second, func() {
			drone.ParseFlightData(myBuffer)
		})

		gobot.After(5*time.Second, func() {
			drone.Land()
		})
		gobot.After(5*time.Second, func() {
			drone.Halt()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

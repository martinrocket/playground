package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

var Drone = tello.NewDriver("8888")
var b = make([]byte, 999)

func main() {

	work := func() {
		Drone.TakeOff()

		gobot.After(5*time.Second, func() {
			Drone.Land()

		})

		gobot.After(6*time.Second, func() {
			data, err := Drone.ParseFlightData(b)
			fmt.Printf("%+v\n%+v\n%s", data, string(b[:999]), err)

		})

		gobot.After(10*time.Second, func() {

			Drone.Halt()
		})

	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{Drone},
		work,
	)

	robot.Start()
}

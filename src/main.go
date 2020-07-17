package main

import (
	"time"
)

const location = "Asia/Tokyo"

func main() {
	// initilize time zone
	initTime()

	// define routes and run app
	r := DefineRoutes()
	r.Run()
}

func initTime() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

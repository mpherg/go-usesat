package main

import (
	"fmt"
	"github.com/joshuaferrara/go-satellite"
	"time"
)

func main() {
	s := satellite.TLEToSat("1 19548U 88091B   17296.36684163 -.00000304  00000-0  00000+0 0  9992", "2 19548  14.4482   6.6503 0035813 306.2967 148.8220  1.00277935 93691", "wgs72")

	// get the current time
	now := time.Now()
	fmt.Println(now)

	ECIpos, ECIvel := satellite.Propagate(s, now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Println(ECIpos)
	fmt.Println(ECIvel)

	gmst := satellite.GSTimeFromDate(now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	altitude, _, latlonrad := satellite.ECIToLLA(ECIpos, gmst)
	latlondeg := satellite.LatLongDeg(latlonrad)
	fmt.Printf("Lat: %f, Lon: %f, Alt: %f\n", latlondeg.Latitude, latlondeg.Longitude, altitude)
}

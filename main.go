package main

import (
	"fmt"
	mjpeg "mjpeg/mjpeg"
)

func main() {
	// Some variables
	host := "0.0.0.0:3000"
	endpoint := "/video/"
	src := 0
	fps := 15

	// Initalize server elements
	webcam := mjpeg.NewCamera(src, fps)
	server := mjpeg.NewMJPEGServer(webcam, host, endpoint)
	fmt.Printf("Server runing on http://%v%v\n", host, endpoint)

	//Execute read/listen loops
	go webcam.Run()
	server.Run()

}

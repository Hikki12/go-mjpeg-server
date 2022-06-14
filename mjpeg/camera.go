package mjpeg

import (
	"time"

	"gocv.io/x/gocv"
)

type Camera struct {
	src     int
	fps     int
	frame   gocv.Mat
	encoded *gocv.NativeByteBuffer
}

func NewCamera(src int, fps int) *Camera {
	return &Camera{src, fps, gocv.NewMat(), &gocv.NativeByteBuffer{}}
}

//
func (c *Camera) GetFrame() gocv.Mat {
	return c.frame
}

func (c *Camera) GetEncodedFrame() *gocv.NativeByteBuffer {
	return c.encoded
}

// Executes the read loop
func (c *Camera) Run() {
	device, _ := gocv.VideoCaptureDevice(c.src)
	delay := time.Duration(1000/c.fps) * time.Millisecond
	for {
		device.Read(&c.frame)
		c.encoded, _ = gocv.IMEncode(".jpg", c.frame)
		time.Sleep(delay)
	}
}

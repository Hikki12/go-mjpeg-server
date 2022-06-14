package mjpeg

import (
	"net/http"
	"sync"
	"time"
)

type MJPEGServer struct {
	camera   *Camera
	host     string
	endpoint string
}

func NewMJPEGServer(camera *Camera, host string, endpoint string) *MJPEGServer {
	return &MJPEGServer{camera, host, endpoint}
}

func (m *MJPEGServer) Run() {
	mutex := &sync.Mutex{}
	path := m.endpoint
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
		data := ""
		for {
			frame := *m.camera.GetEncodedFrame()
			mutex.Lock()
			data = "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(frame.GetBytes()) + "\r\n\r\n"
			mutex.Unlock()
			time.Sleep(15 * time.Millisecond)
			w.Write([]byte(data))
		}
	})
	http.ListenAndServe(m.host, nil)
}

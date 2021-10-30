package utility

import (
	"net"
)

var control net.UDPConn

func InitCarControl() bool {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 1324,
	})
	control = *socket
	return !IsErr(err, "UDP Control Init Failed")

}

func SendCarControl(buf string) bool {
	sendControl := []byte(buf)
	_, err := control.Write(sendControl)
	return !IsErr(err, "UDP Control Send Failed")
}

func CloesCarControl() {
	control.Close()
}

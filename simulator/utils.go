package simulator

import "net"

func MustParseMacAddr(mac string) net.HardwareAddr {
	retval, err := net.ParseMAC(mac)
	if err != nil {
		panic(err)
	}
	return retval
}

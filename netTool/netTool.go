package netTool

import (
	"net"
	"strings"
)

var HostIp = GetIp()
var ServerName = GetIp()

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		ip := strings.Split(addr.String(), "/")[0]
		code := strings.Split(ip, ".")
		switch code[0] {
		case "10", "127":
			continue
		default:
			return ip
		}
	}
	panic(addrs)
}

func GetIpByUrl(url string) ([]net.IP, error) {
	ns, err := net.LookupIP(url)
	return ns, err

}

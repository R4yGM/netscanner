package netmask

import (
	"fmt"
	"log"
	"net"
)

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, nil

	default:
	return ips[1 : len(ips)-1], nil
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	tmp, err := Hosts("127.0.0.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmp)

	tmp, err = Hosts("192.168.1.1/32") //panic here
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmp)
}

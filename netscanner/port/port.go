
package port

import (
	"net"
	"strconv"
	"time"
	netserv"./netserv"
)

type ScanResult struct {		
	Port    string
	State   string
	Service string
}

func ScanPort(protocol, hostname string, port int) ScanResult {

				var proto *netserv.Protoent
         var serv *netserv.Servent
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol+", "}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed, "
	}else {
		defer conn.Close()
				result.State = "Open, "
	}
	proto = netserv.GetProtoByName(protocol)
	serv = netserv.GetServByPort(port, proto)

	if serv != nil {
		result.Service = serv.Name + "\n"
	}
	if result.Service == ""{
		result.Service = "\n"
	}

	return result
}
func ScanFromTo(hostname string, protocol string, startport int, endport int) []ScanResult {

	var results []ScanResult
	 for i := startport; i <= endport; i++ {
		results = append(results, ScanPort(protocol, hostname, i))
	}


	return results

}
 func InitialScan(hostname string, protocol string) []ScanResult {

	var results []ScanResult

	 for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort(protocol, hostname, i))
	}
	return results
	}

func CompleteScan(hostname string, protocol string) []ScanResult {

 	var results []ScanResult

 	 for i := 0; i <= 65535; i++ {
 		results = append(results, ScanPort(protocol, hostname, i))
 	}


	return results
}

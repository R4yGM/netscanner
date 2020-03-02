
/*  _____     ____
 / ___/__  / __/______ ____
/ (_ / _ \_\ \/ __/ _ `/ _ \
\___/\___/___/\__/\_,_/_//_/
                            */


package port

import (
	"net"
	"strconv"
	"time"
	"honnef.co/go/netdb"
)

type ScanResult struct {
	Port    string
	State   string
	Service string
}

func ScanPort(protocol, hostname string, port int) ScanResult {

				var proto *netdb.Protoent
         var serv *netdb.Servent
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol+", "}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed, "
	}else {
		defer conn.Close()
				result.State = "Open, "
	}
	proto = netdb.GetProtoByName(protocol)
	serv = netdb.GetServByPort(port, proto)

	if serv != nil {
		result.Service = serv.Name + "\n"
	}
	if result.Service == ""{
		result.Service = "\n"
	}

	return result
}

 func InitialScan(hostname string, protocol string) []ScanResult {

	var results []ScanResult

	 for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort(protocol, hostname, i))
	}


	return results
}

func WideScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}

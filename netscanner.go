package netscanner

import (
	"fmt"
  "strings"
	"time"
	PortScan"github.com/R4yGM/netscanner/port"
	header"github.com/R4yGM/netscanner/header"
	Info"github.com/R4yGM/netscanner/info"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1{
		Help()
		return
	}

	arg := os.Args[1]


	if arg == "completescan" || arg == "cs"{
		if len(os.Args) <= 4{
			fmt.Println("Missing parameters! \nUsage : cs,completescan <protocol> <hostname> <savefile>(ex. initscan udp 127.0.0.1 true)\n-protocol - can be tcp or udp\n-hostname - hostname of the target - localhost/127.0.0.1\n-savefile - bool that saves the scan - true / false \ntype help if you need help")
			return
		}
		var proto = os.Args[2]
		var hostname = os.Args[3]
		var booll = os.Args[4]
		if booll != "false" || booll != "true"{
			//fmt.Println("wrong arguments : \nUsage : cs,completescan <protocol> <hostname> <savefile>(ex. initscan udp 127.0.0.1 true)\n-protocol - can be tcp or udp\n-hostname - hostname of the target - localhost/127.0.0.1\n-savefile - bool that saves the scan - true / false \ntype help if you need help")
			//return
			booll = "false"
		}
		booll1, err := strconv.ParseBool(booll)
		if err != nil{
			fmt.Println(err)
		}

		CompleteScan(proto, hostname, booll1)
		return
	}
	if arg == "help" || arg == "h"{
		Help()
		return
	}

	if arg == "scanfromto" || arg == "sft"{									
		if len(os.Args) <= 6{				
			fmt.Println("\nMissing parameters! \nUsage : sft,scanfromto  <protocol> <hostname> <startport> <endport> <savefile> (ex. scanfromto tcp 127.0.0.1 20 100 false)\n-protocol - can be tcp or udp\n-hostname - hostname of the target - localhost/127.0.0.1\n-startport - first port to scan\n-endport - last port to scan\n-savefile - bool that saves the scan - true / false\ntype help if you need more help\n")
			return
		}
		var proto = os.Args[2]
		var hostname = os.Args[3]
		var	startport = os.Args[4]
		var	endport = os.Args[5]
		var	booll = os.Args[6]
		booll1, err := strconv.ParseBool(booll)
		if err != nil{
			fmt.Println(err)
		}
		sp, err := strconv.Atoi(startport)
	if err != nil {
		fmt.Println(err)
	}

	ep, err := strconv.Atoi(endport)
if err != nil {
	fmt.Println(err)
}
		ScanFromTo(proto,hostname,sp,ep,booll1)
		return
	}
	if arg == "information" || arg == "info"{
		var info = Infor()
		fmt.Println(info)
		return
	}
	if arg == "version" || arg == "v"{
		var ver = Version()
		fmt.Println(ver)
		return
	}


	if arg == "scanport" || arg == "sp"{
			if len(os.Args) <= 4{
			fmt.Println("Missing parameters! \nUsage : sp,scanport  protocol hostname port (ex. scanport tcp 127.0.0.1 80)\n-protocol - can be tcp or udp\n-hostname - hostname of the target - localhost/127.0.0.1 -port - port to scan\ntype help if you need help")
			return
		}

		var proto = os.Args[2]
		var hostname = os.Args[3]
		var	port = os.Args[4]
		porto, err := strconv.Atoi(port)
		if err != nil{
			fmt.Println(err)
		}
		ScanPort(proto, hostname, porto)
		return
	}

	if arg == "initscan" || arg == "in"{
		if len(os.Args) <= 4{
			fmt.Println("Missing parameters! \nUsage : in,initscan <protocol> <hostname> <savefile>(ex. initscan udp 127.0.0.1 false)\n-protocol - can be tcp or udp\n-hostname - hostname of the target - localhost/127.0.0.1\n-savefile - bool that saves the scan\ntype help if you need help")
			return
		}
		var proto = os.Args[2]
		var hostname = os.Args[3]
		var booll = os.Args[4]

		booll1, err := strconv.ParseBool(booll)
		if err != nil{
			fmt.Println(err)
		}

		InitialScan(proto, hostname, booll1)
		return
	}
	if arg == "help" || arg == "h"{
		Help()
		return
	}
	fmt.Println("Unkown command\n")
		Help()
		return
}
func ScanPort(protocol string, hostname string, port int)(texe string){

	var tex = header.AsciiTitle()
	fmt.Println(tex)
	dtStart := time.Now()
	fmt.Println("Port Scanning\n")
	fmt.Println("  Port      status   service\n")
	results := PortScan.ScanPort(protocol,hostname,port)
	var replacer = strings.NewReplacer("{", " ", "}", "")


var str = replacer.Replace(fmt.Sprint(results))
	fmt.Println(strings.Trim(fmt.Sprint(str), "[]"))
	dtEnd := time.Now()
fmt.Println("Scan started at :", dtStart.String())
fmt.Println("And finished at :", dtEnd.String())
fmt.Println("\n")
return str
}
func InitialScan(protocol string, hostname string, savefile bool)(texe string){

	var tex = header.AsciiTitle()
	fmt.Println(tex)
	dtStart := time.Now()
	fmt.Println("Port Scanning\n")
	fmt.Println("  Port      status   service\n")
	results := PortScan.InitialScan(protocol,hostname)
	var replacer = strings.NewReplacer("{", " ", "}", "")


var str = replacer.Replace(fmt.Sprint(results))
if savefile == true{
	f, err := os.Create("CompleteScan"+dtStart.String()+ ".txt")
		if err != nil {
				fmt.Println(err)
				return
		}


var str1 = replacer.Replace(fmt.Sprint(results))
f.WriteString(strings.Trim(fmt.Sprint(str1), "[]"))

f.Close()
}
	fmt.Println(strings.Trim(fmt.Sprint(str), "[]"))
	dtEnd := time.Now()
fmt.Println("Scan started at :", dtStart.String())
fmt.Println("And finished at :", dtEnd.String())
if savefile == true{
fmt.Println("Scan saved in CompleteScan"+dtStart.String()+ ".txt")
}
fmt.Println("\n")
return str
}
func CompleteScan(protocol string, hostname string, savefile bool)(texe string){

	var tex = header.AsciiTitle()
	fmt.Println(tex)
	dtStart := time.Now()
	fmt.Println("Port Scanning\n")
	fmt.Println("  Port      status   service\n")
	results := PortScan.CompleteScan(protocol,hostname)
	var replacer = strings.NewReplacer("{", " ", "}", "")
	var str = replacer.Replace(fmt.Sprint(results))
if savefile == true{
	f, err := os.Create("CompleteScan"+dtStart.String()+ ".txt")
		if err != nil {
				fmt.Println(err)
				return
		}


var str1 = replacer.Replace(fmt.Sprint(results))
f.WriteString(strings.Trim(fmt.Sprint(str1), "[]"))

f.Close()
}

	fmt.Println(strings.Trim(fmt.Sprint(str), "[]"))
	dtEnd := time.Now()
fmt.Println("Scan started at :", dtStart.String())
fmt.Println("And finished at :", dtEnd.String())
if savefile == true{

fmt.Println("Scan saved in CompleteScan"+dtStart.String()+ ".txt")
}


fmt.Println("\n")

return str
}
func ScanFromTo(hostname string, protocol string, startport int, endport int,savefile bool)(texe string){

		if startport > endport {
			var initP = startport
			startport = endport
			endport = initP
		}

		var tex = header.AsciiTitle()
		fmt.Println(tex)
		dtStart := time.Now()
		fmt.Println("Port Scanning\n")
		fmt.Println("  Port      status   service\n")
		results := PortScan.ScanFromTo(protocol,hostname,startport,endport)
		var replacer = strings.NewReplacer("{", " ", "}", "")
		if savefile == true{
			f, err := os.Create("CompleteScan"+dtStart.String()+ ".txt")
				if err != nil {
						fmt.Println(err)
						return
				}


		var str1 = replacer.Replace(fmt.Sprint(results))
		f.WriteString(strings.Trim(fmt.Sprint(str1), "[]"))

		f.Close()
		}


	var str = replacer.Replace(fmt.Sprint(results))
		fmt.Println(strings.Trim(fmt.Sprint(str), "[]"))
		dtEnd := time.Now()
	fmt.Println("Scan started at :", dtStart.String())
	fmt.Println("And finished at :", dtEnd.String())
	if savefile == true{

	fmt.Println("Scan saved in CompleteScan"+dtStart.String()+ ".txt")
	}
	fmt.Println("\n")
	return str

}

func Infor()(tex string){
	var ver =Info.Version()
	var own = Info.Owner()
	var gith = Info.Github()
  tex = "\nversion : "+string(ver)+"\nmade by :"+own+"\nGithub Repository : "+gith+"\n"
	return
}
func Version()(tex string){
	var ver =Info.Version()
	tex = ver
	return
}

func Help()(tex string) {
	var helpe = Info.Helper()
	tex = helpe
	fmt.Println(tex)
	return
}

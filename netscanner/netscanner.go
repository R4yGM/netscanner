package main

import (
	"fmt"
  "strings"
	"time"
	PortScan"./port"
	AsciiGen"./header"
	Info"./info"
	"os"
	"strconv"
)

func main() {
  //SCAN SINGUAL PortS

	arg := os.Args[1]

	if arg == "scanport"{
		var proto = os.Args[2]
		var hostname = os.Args[3]
		var	port = os.Args[4]
		//port = int(port)
		porto, err := strconv.Atoi(port)
		if err != nil{
			fmt.Println(err)
		}
		ScanPort(proto, hostname, porto)
	}
	/*var tex = AsciiGen.AsciiTitle()
	fmt.Println(tex)
	dtStart := time.Now()
	fmt.Println("Port Scanning")
  fmt.Println("  Port      status   service\n")
	//results := PortScan.InitialScan("127.0.0.1","tcp")
  results := PortScan.ScanPort("tcp","127.0.0.1",8080)
  var replacer = strings.NewReplacer("{", " ", "}", "")


var str = replacer.Replace(fmt.Sprint(results))
  fmt.Println(strings.Trim(fmt.Sprint(str), "[]"))
	dtEnd := time.Now()
fmt.Println("Scan started at :", dtStart.String())
fmt.Println("And finished at :", dtEnd.String())*/
}
func ScanPort(protocol string, hostname string, port int){

	var tex = AsciiGen.AsciiTitle()
	fmt.Println(tex)
	dtStart := time.Now()
	fmt.Println("Port Scanning\n")
	fmt.Println("  Port      status   service\n")
	results := PortScan.ScanPort(protocol,hostname,port)  //udp tcp
	var replacer = strings.NewReplacer("{", " ", "}", "")


var str = replacer.Replace(fmt.Sprint(results))
	fmt.Println(strings.Trim(fmt.Sprint(str), "[]"))
	dtEnd := time.Now()
fmt.Println("Scan started at :", dtStart.String())
fmt.Println("And finished at :", dtEnd.String())
}

func info()(tex string){
	var ver =Info.Version
	var own = Info.Owner
	var gith = Info.Github
	//"version : "+string(ver)+"	"+
//	var ret = "made by :"+own+"\n"+"Github Repository : "+gith
	fmt.Printf("version : ",ver)
	fmt.Printf("	made by :",own)
	fmt.Printf("	Github Repository : ",gith)

//	fmt.Println(ret)
	return
}
func version(){
	var ver =Info.Version
	fmt.Println(ver)
	return
}

func help() {
	var help = Info.Helper
	fmt.Println(help)
}

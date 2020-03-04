package main

import(
	
	netscanner"./netscanner"

)

func main(){

	netscanner.CompleteScan("tcp","127.0.0.1", true)
}

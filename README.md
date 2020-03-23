[![Build Status](https://travis-ci.org/R4yGM/netscanner.svg?branch=master)](https://travis-ci.org/R4yGM/netscanner)
# netscanner
netscanner - TCP/UDP scanner to find open or closed ports
## installation 
you have to run this command to install the program
```shell
$ go get github.com/R4yGM/netscanner
```

## Usage
### Commands :
```shell
netscanner 0.1 : TCP/UDP port scanner

Usage: netscanner <command>

Commands:

   sp,scanport - scan a singular port: usage sp,scanport <protocol> <hostname> <port> (ex. scanport tcp 127.0.01 80)
   in,initscan - scan the first 1023 ports: usage in,initscan <protocol> <hostname> <savefilebool>(ex. initscan udp 127.0.0.1 false)
   cs,completescan - scan all the ports: usage cs,completescan <protocol> <hostname> <savefilebool>(ex. completescan tcp 127.0.0.1 true)
   sft,scanfromto - scan the ports from a port given to another port given: usage sft,scanfromto <protocol> <hostname> <startport> <endport> <savefile>(ex. sft tcp 127.0.0.1 80 443 true)
   h,help - shows this text
   v,version - show the current version of the program
   info,information - shows some information about the program
   msp,maskscanport - scan a port of all the IPs in Subnet: usage msp,maskscanport <protocol> <hostname> <port> <savefilebool>(ex. msp tcp 127.0.0.1/24 true)
   
``` 
### Parameters:

```
<protocol> - protocol to use - can be tcp or udp
<hostname> - hostname of the target - ex. 127.0.0.1
<port> - port to scan
<savefile> - bool that saves the scan in a file - true / false
<startport> - port to start scanning
<endport> - last port to scan
<hostname>(for mask scan) - must contain /24 or /32
```
### Example

```
$ netscanner sp tcp 127.0.0.1 80

 Port     status   service

 80/tcp,  Closed,  http

Scan started at : 2020-03-07 12:16:08.4845743 +0100 CET m=+0.002086201
And finished at : 2020-03-07 12:16:08.4854596 +0100 CET m=+0.002971501

```


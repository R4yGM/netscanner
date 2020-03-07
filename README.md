# netscanner
netscanner - TCP/UDP scanner to find open or closed ports
## installation 
you have to run this command to installer the installer
```shell
$ wget https://gist.githubusercontent.com/R4yGM/b989427de2c74d04e2ed1b4257108e0c/raw/5c3a36a107e36d1a2798fce4a330f0700bb2f2de/installer.sh
```
then you will have to make executable installer.sh
```shell
$ chmod +x netscanner.sh
```
and then run the installer by typing
```shell
$ sudo ./netscanner.sh
```
## Usage
### Commands :
```shell
netscanner 0.1

Usage: netscanner <command>

Commands:

   sp,scanport - scan a singular port: usage sp,scanport <protocol> <hostname> <port> (ex. scanport tcp 127.0.0.1 80)
   in,initscan - scan the first 1023 ports: usage in,initscan <protocol> <hostname> <savefilebool>(ex. initscan udp 127.0.0.1 false)
   cs,completescan - scan all the ports: usage cs,completescan <protocol> <hostname> <savefilebool>(ex. completescan tcp 127.0.0.1 true)
   sft,scanfromto - scann the ports from a port given to another port given: usage sft,scanfromto <protocol> <hostname> <startport> <endport> <savefile>(ex. sft tcp 127.0.0.1 80 443 true)
``` 
### Parameters:

```
<protocol> - protocol to use - can be tcp or udp
<hostname> - hostname of the target - ex. 127.0.0.1
<port> - port to scan
<savefile> - bool that saves the scan in a file - true / false
<startport> - port to start scanning
<endport> - last port to scan
```
### Example

```
$ netscanner sp tcp 127.0.0.1 80

 Port     status   service

 80/tcp,  Closed,  http

Scan started at : 2020-03-07 12:16:08.4845743 +0100 CET m=+0.002086201
And finished at : 2020-03-07 12:16:08.4854596 +0100 CET m=+0.002971501

```

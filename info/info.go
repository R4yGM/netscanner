package info

import (
//  "io/ioutil"
//  "encoding/json"
  //"fmt"
//  "os"
)
var ver = "0.1"
var owner = "R4yan"
var github = "https://github.com/R4yGM/netscanner"
var help = "\nnetscanner 0.1\n     \nUsage: netscanner <command>\n\nCommands:\n\n   sp,scanport - scan a singular port: usage sp,scanport <protocol> <hostname> <port> (ex. scanport tcp 127.0.01 80)\n   in,initscan - scan the first 1023 ports: usage in,initscan <protocol> <hostname> <savefilebool>(ex. initscan udp 127.0.0.1 false)\n   cs,completescan - scan all the ports: usage cs,completescan <protocol> <hostname> <savefilebool>(ex. completescan tcp 127.0.0.1 true)\n   sft,scanfromto - scann the ports from a port given to another port given: usage sft,scanfromto <protocol> <hostname> <startport> <endport> <savefile>(ex. sft tcp 127.0.0.1 80 443 true)\n   h,help - shows this text\n   v,version - show the current version of the program\n   info,information - shows some information about the program\n"

type Info struct {
	Version string `json:"version"`
	Owner   string    `json:"owner"`
  Github string `json:"github"`
  Help string  `json:"help"`
}

func Version()(tex string){
//  file, _ := ioutil.ReadFile("info.json")

/*  	data := Info{}

  	_ = json.Unmarshal([]byte(file), &data)*/
    //fmt.Println(ver)
    tex = ver
    return tex
}
func Owner()(tex string){

//  file, _ := ioutil.ReadFile("info.json")

  /*  data := Info{}

    _ = json.Unmarshal([]byte(file), &data)*/
      //  fmt.Println(owner)
    tex = owner
    return tex

}
func Github()(tex string){

//  file, _ := ioutil.ReadFile("./info.json")

    /*data := Info{}

    _ = json.Unmarshal([]byte(file), &data)
    tex = data.Github*/
    tex = github
    //fmt.Println(tex)
    return tex

}
func Helper()(tex string){

  //jsonFile, err := os.Open("github.com/R4yGM/netscanner/data.json")
/*
     if err != nil {
         fmt.Println(err)
     }

     defer jsonFile.Close()

     var in Info


      responseData, err := ioutil.ReadAll(jsonFile)
      json.Unmarshal(responseData, &in)
     fmt.Println(in.Help)*/
     tex = help
     return

}

package info

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
  "os"
)
type Info struct {
	Version int `json:"version"`
	Owner   string    `json:"owner"`
  Github string `json:"github"`
  Help string  `json:"help"`
}

func Version()(tex int){
  file, _ := ioutil.ReadFile("info.json")

  	data := Info{}

  	_ = json.Unmarshal([]byte(file), &data)
    tex = data.Version
    return tex
}
func Owner()(tex string){

  file, _ := ioutil.ReadFile("info.json")

    data := Info{}

    _ = json.Unmarshal([]byte(file), &data)
    tex = data.Owner
    return tex

}
func Github()(tex string){

  file, _ := ioutil.ReadFile("./info.json")

    data := Info{}

    _ = json.Unmarshal([]byte(file), &data)
    tex = data.Github
    return tex

}
func Helper()(tex string){

  jsonFile, err := os.Open("data.json")

     if err != nil {
         fmt.Println(err)
     }

     defer jsonFile.Close()

     var in Info


      responseData, err := ioutil.ReadAll(jsonFile)
      json.Unmarshal(responseData, &in)
     fmt.Println(in.Help)
     return

}

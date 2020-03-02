package info

import (
  "io/ioutil"
  "encoding/json"
)
type Info struct {
	version int `json: "version"`
	owner   string    `json: "owner"`
  help string  `json: "help"`
  github string `json: "github"`
}

func Version()(tex int){
  file, _ := ioutil.ReadFile("info.json")

  	data := Info{}

  	_ = json.Unmarshal([]byte(file), &data)
    tex = data.version
    return tex
}
func Owner()(tex string){

  file, _ := ioutil.ReadFile("info.json")

    data := Info{}

    _ = json.Unmarshal([]byte(file), &data)
    tex = data.owner
    return tex

}
func Github()(tex string){

  file, _ := ioutil.ReadFile("info.json")

    data := Info{}

    _ = json.Unmarshal([]byte(file), &data)
    tex = data.github
    return tex

}
func Helper()(tex string){

  file, _ := ioutil.ReadFile("info.json")

    data := Info{}

    _ = json.Unmarshal([]byte(file), &data)
    tex = data.help
    return tex

}

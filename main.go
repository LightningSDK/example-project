package main

import (
	"encoding/xml"
	"fmt"
	"github.com/lightningsdk/ui"
	"github.com/lightningsdk/ui/standard"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func main() {
	fc, _ := ioutil.ReadFile("config/content/content-management/home.yaml")
	_ = ui.Init()
	r := &standard.Page{}
	_ = yaml.Unmarshal(fc, r)
	fmt.Println(r.Render())
	x, _ := xml.Marshal(r)
	fmt.Println(string(x))

	i := []interface{}{}
	xm := xml.Unmarshal([]byte("<div>something</div>"), &i)
	fmt.Println(xm)
}

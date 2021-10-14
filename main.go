package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/walkersumida/jadtbuilder"
)

// Item is Alfred's item struct.
type Item struct {
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Menu is Alfred's menu struct.
type Menu struct {
	Items []Item `json:"items"`
}

func outputFormat(item Item) {
	var menu Menu
	item.Icon = "./icon.png"
	menu.Items = append(menu.Items, item)

	menuJSON, _ := json.Marshal(menu)
	fmt.Println(string(menuJSON))
}

func main() {
	var item Item

	flag.Parse()
	inputDt := strings.Join(flag.Args(), " ")

	builtDt := jadtbuilder.Build(inputDt)

	item.Title = builtDt
	item.Arg = builtDt

	outputFormat(item)
}

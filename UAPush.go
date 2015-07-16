package main

import (
"time"
"encoding/json"
"fmt"

"github.com/jamesaanderson/go-urbanairship"

//"log"
)

type Push struct {
  Audience     interface{}  `json:"audience"`
  Notification Notification `json:"notification"`
  DeviceTypes  interface{}  `json:"device_types"`
}

type Notification struct {
  Alert string `json:"alert"`
}

func main() {
  fmt.Println(time.Date)
  c := urbanairship.NewClient(nil, "aMF4HdJdTFCbFMKhnh5mgA", "DiWlAF22QrGzRN8Qo_Shyw")

  type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
  }
  group := &ColorGroup{
    ID:     1,
    Name:   "Reds",
    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
  }

  a, _ := json.Marshal(group)
  fmt.Println(string(a))

  message := &Push{
    Audience: { "alias" : "1" },
    Notification: &Alert{ "yo" },
    DeviceTypes: { ["ios"] }
  }
  a, _ := json.Marshal(message)
  fmt.Println(string(a))
  //c.Push(message)
}


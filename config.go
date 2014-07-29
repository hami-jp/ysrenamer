package ysrenamer

import (
  "fmt"
  "os"
  "io/ioutil"
  "runtime"

  . "github.com/bitly/go-simplejson"
)

var ConfigPath string = func() string {
  switch runtime.GOOS {
    case "windows":
      return os.Getenv("APPDATA") + "\\ysrenamer\\ysrenamer.json"
    default:
      return os.Getenv("HOME") + "/.ysrenamer.json"
  }
}()

var Config *Json = func() *Json {
  var raw_data []byte
  var err error
  var js *Json

  _, err = os.Stat(ConfigPath)
  if os.IsNotExist(err) {
    raw_data = []byte("{}")
    err = ioutil.WriteFile(ConfigPath, raw_data, 0644)
  } else {
    raw_data, err = ioutil.ReadFile(ConfigPath)
  }
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  js, err = NewJson(raw_data)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  return js
}()

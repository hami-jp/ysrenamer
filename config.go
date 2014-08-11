package ysrenamer

import (
  "fmt"
  "os"
  "io/ioutil"
  "runtime"

  . "github.com/bitly/go-simplejson"
  "code.google.com/p/mahonia"
)

var ConfigPath string = func() string {
  switch runtime.GOOS {
    case "windows":
      return os.Getenv("APPDATA") + "\\ysrenamer\\ysrenamer.json"
    default:
      return os.Getenv("HOME") + "/.ysrenamer.json"
  }
}()

var Config *Json = readConfigFile()

func readConfigFile() *Json {
  raw_data := readFile(ConfigPath)
  js, err := NewJson(raw_data)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  return js
}

func readFile(path string) []byte {
  var data []byte

  _, err := os.Stat(path)
  if os.IsNotExist(err) {
    data = []byte("{}")
    err = ioutil.WriteFile(path, data, 0644)
  } else {
    data, err = ioutil.ReadFile(path)
  }

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  data = []byte(convertStringToUtf8(string(data)))

  return data
}

func convertStringToUtf8(str string) string {
  return mahonia.NewDecoder("utf-8").ConvertString(str)
}

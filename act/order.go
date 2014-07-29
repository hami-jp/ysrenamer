package act

import (
  "fmt"
  "os"

  "github.com/yuukichi/ysrenamer"
)

func ProcOrder(command string, files []string) []string {
  processed := make([]string, len(files))

  conf := ysrenamer.Config.Get(command) 
  action, _ := conf.Get("action").String()
  args, _ := conf.Get("args").StringArray()

  switch action {
    case "replace":
      processed = Replace(files, args) 
    case "regexp_replace":
      processed = RegexpReplace(files, args)
    case "order":
      processed = Order(files, args)
    default:
      fmt.Println(command + " command not found.")
      os.Exit(1)
  }

  return processed
}

func Order(files []string, commands []string) []string {
  processed := files 
  for _, v := range commands {
    processed = ProcOrder(v, processed)
  }
  return processed
}

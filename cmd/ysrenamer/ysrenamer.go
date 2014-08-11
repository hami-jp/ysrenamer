package main

import (
  "fmt"
  "os"

  "github.com/yuukichi/ysrenamer"
)

func main() {
  initialize(os.Args)
}

func initialize(args []string) {
  checkArgumentsCount(args)
  ysrenamer.Rename(args[2:], args[1]) 
}

func checkArgumentsCount(args []string) {
  if len(args) < 3 {
    fmt.Println("usage: " + args[0] + " command args")
    os.Exit(1)
  }
}

package main

import (
  "fmt"
  "os"

  "github.com/yuukichi/ysrenamer/act"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("usage: " + os.Args[0] + " command args")
    os.Exit(1)
  }

  files := os.Args[2:]
  dests := act.ProcOrder(os.Args[1], files)
 
  var dest string
  for i, src := range files {
    dest = dests[i]
    if err := os.Rename(src, dest) ; err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
    fmt.Println(src + " -> " + dest)
  }
}

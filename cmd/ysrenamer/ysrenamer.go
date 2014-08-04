package main

import (
  "fmt"
  "os"
  "path/filepath"

  "code.google.com/p/go.text/unicode/norm"
  "github.com/yuukichi/ysrenamer/act"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("usage: " + os.Args[0] + " command args")
    os.Exit(1)
  }

  files := os.Args[2:]
  files_encoded := make([]string, len(files))
  for k, v := range files {
    output := DenormalizeUnicodeString(v)
    files_encoded[k] = output
  }

  dests := act.ProcOrder(os.Args[1], files_encoded)
 
  for i, src := range files {
    dest := dests[i]
    dir := filepath.Dir(dest)
    if _, err := os.Stat(dir) ; os.IsNotExist(err) {
      if err := os.MkdirAll(dir, 0755) ; err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }
    }

    if err := os.Rename(src, dest) ; err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
    fmt.Println(src + " -> " + dest)
  }
}

func DenormalizeUnicodeString(str string) string {
  buf := []byte(str)
  buf = norm.NFC.Bytes(buf)
  return string(buf)
}

package ysrenamer

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "path/filepath"

  "code.google.com/p/go.text/unicode/norm"
)

func Rename(srcs []string, action string) {
  dests := ProcOrder(action, DenormalizeStringArray(srcs))
  renameFollowingArray(srcs, dests)
}

func renameFollowingArray(srcs []string, dests []string) {
  for i, src := range srcs {
    dest := dests[i]
    dir := filepath.Dir(dest)
    MakeDirectoryIfIsNotExist(dir)
    rename(src, dest)
  }
}

func rename(src string, dest string) {
  if err := move(src, dest) ; err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  fmt.Println(src + " -> " + dest)
}

func move(src string, dest string) error {
  var cmd *exec.Cmd
  if runtime.GOOS != "windows" {
    cmd = exec.Command("mv", src, dest)
  } else {
    cmd = exec.Command("cmd", "/C", "move", src, dest)
  }
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func MakeDirectoryIfIsNotExist(dir string) {
  if _, err := os.Stat(dir) ; os.IsNotExist(err) {
    if err := os.MkdirAll(dir, 0755) ; err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
  }
}

func DenormalizeStringArray(arr []string) []string {
  new_arr := make([]string, len(arr))
  for k, v := range arr {
    new_arr[k] = DenormalizeUnicodeString(v)
  }
  return new_arr 
}

func DenormalizeUnicodeString(str string) string {
  buf := []byte(str)
  buf = norm.NFC.Bytes(buf)
  return string(buf)
}

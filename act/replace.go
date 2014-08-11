package act

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "path/filepath"
)

func procReplace(path string, options []string) string {
  CheckOptionsCount(options, 2)
  
  n := getReplaceCount(options) 

  dir := filepath.Dir(path)
  name, ext := GetNameAndExtensionIfOptionEnabled(path, options, 3)

  return JoinFilePath(dir, strings.Replace(name, options[0], options[1], n) + ext)  
}

func getReplaceCount(options []string) int {
  if len(options) > 2 {
    if i, _ := strconv.ParseBool(options[2]) ; i {
      n, err := strconv.Atoi(options[2])
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }
      return n
    }
  }
  return 1
}

func Replace(files []string, options []string) []string {
  processed := make([]string, len(files))
  for i, v := range files {
    processed[i] = procReplace(v, options)
  }
  return processed
}

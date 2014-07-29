package act

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "path/filepath"
)

func proc_replace(file string, options []string) string {
  var err error
  if len(options) < 2 {
    fmt.Println("error: this action has invalid arguments. please correct your config file.")
    os.Exit(1)
  }
  
  n := 1

  if len(options) > 2 {
    if i, _ := strconv.ParseBool(options[2]) ; i {
      n, err = strconv.Atoi(options[2])
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }
    }
  }

  var suffix string
  base_name := filepath.Base(file)
  dir := filepath.Dir(file)

  if len(options) > 3 {
    if i, _ := strconv.ParseBool(options[3]) ; i {
      suffix = filepath.Ext(base_name)
      base_name = base_name[0:len(base_name)-len(suffix)]
    }
  }

  return filepath.Join(dir, strings.Replace(base_name, options[0], options[1], n) + suffix)  
}

func Replace(files []string, options []string) []string {
  processed := make([]string, len(files))
  for i, v := range files {
    processed[i] = proc_replace(v, options)
  }
  return processed
}

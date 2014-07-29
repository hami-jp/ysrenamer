package act

import (
  "fmt"
  "os"
  "regexp"
  "strconv"
  "path/filepath"
)

func proc_regexp_replace(file string, options []string) string {
  if len(options) < 2 {
    fmt.Println("error: this action has invalid arguments. please correct your config file.")
    os.Exit(1)
  }

  var suffix string
  base_name := filepath.Base(file)
  dir := filepath.Dir(file)

  if len(options) > 2 {
    if i, _ := strconv.ParseBool(options[2]) ; i {
      suffix = filepath.Ext(base_name)
      base_name = base_name[0:len(base_name)-len(suffix)]
    }
  }
  
  reg := regexp.MustCompile(options[0])
  return filepath.Join(dir, reg.ReplaceAllString(base_name, options[1]) + suffix)
}

func RegexpReplace(files []string, options []string) []string {
  processed := make([]string, len(files))
  for i, v := range files {
    processed[i] = proc_regexp_replace(v, options)
  }
  return processed
}

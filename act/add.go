package act

import (
  "fmt"
  "os"
  "strconv"
  "path/filepath"
)

func add_string(base string, add string, rear bool) string {
  if rear {
    return base + add
  } else {
    return add + base
  }
}

func proc_add(file string, options []string, rear bool) string {
  if len(options) < 1 {
    fmt.Println("error: this action has invalid arguments. please correct your config file.")
    os.Exit(1)
  }

  var suffix string
  base_name := filepath.Base(file)
  dir := filepath.Dir(file)

  if len(options) > 1 {
    if i, _ := strconv.ParseBool(options[1]) ; i {
      suffix = filepath.Ext(base_name)
      base_name = base_name[0:len(base_name)-len(suffix)]
    }
  }
  
  return filepath.Join(dir, add_string(base_name, options[0], rear) + suffix)
}

func Add(files []string, options []string, rear bool) []string {
  processed := make([]string, len(files))
  for i, v := range files {
    processed[i] = proc_add(v, options, rear)
  }
  return processed
}

func AddFront(files []string, options []string) []string {
  return Add(files, options, false)
}

func AddRear(files []string, options []string) []string {
  return Add(files, options, true) 
}

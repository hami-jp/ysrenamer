package act

import (
  "fmt"
  "os"
  "strconv"
  "runtime"
  "path/filepath"
)

func CheckOptionsCount(options []string, count int) {
  if len(options) < count {
    fmt.Println("error: this action has invalid arguments. please correct your config file.")
    os.Exit(1)
  }
}

func GetNameAndExtensionIfOptionEnabled(path string, options []string, optionIndex int) (string, string) {
  if len(options) > optionIndex {
    if i, _ := strconv.ParseBool(options[optionIndex]) ; i {
      return GetNameAndExtension(path)
    }
  }
  return filepath.Base(path), "" 
}

func GetNameAndExtension(path string) (string, string) {
  base_name := filepath.Base(path)
  ext := filepath.Ext(base_name)
  name := base_name[0:len(base_name)-len(ext)]
  return name, ext 
}

func JoinFilePath(dir string, name string) string {
  if (runtime.GOOS != "windows" && name[0] == 0x2f) || (runtime.GOOS == "windows" && name[1] == 0x3a) {
    return name
  } else {
    return filepath.Join(dir, name)
  }
}

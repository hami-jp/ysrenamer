package act

import (
  "path/filepath"
)

func addString(base string, add string, rear bool) string {
  if rear {
    return base + add
  } else {
    return add + base
  }
}

func procAdd(path string, options []string, rear bool) string {
  CheckOptionsCount(options, 1)

  dir := filepath.Dir(path)
  name, ext := GetNameAndExtensionIfOptionEnabled(path, options, 1)
  
  return JoinFilePath(dir, addString(name, options[0], rear) + ext)
}


func Add(files []string, options []string, rear bool) []string {
  processed := make([]string, len(files))
  for i, v := range files {
    processed[i] = procAdd(v, options, rear)
  }
  return processed
}

func AddFront(files []string, options []string) []string {
  return Add(files, options, false)
}

func AddRear(files []string, options []string) []string {
  return Add(files, options, true) 
}

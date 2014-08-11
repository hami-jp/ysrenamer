package act

import (
  "regexp"
  "path/filepath"
)

func procRegexpReplace(path string, options []string) string {
  CheckOptionsCount(options, 2)

  dir := filepath.Dir(path)
  name, ext := GetNameAndExtensionIfOptionEnabled(path, options, 2)
  
  reg := regexp.MustCompile(options[0])

  return JoinFilePath(dir, reg.ReplaceAllString(name, options[1]) + ext)
}

func RegexpReplace(files []string, options []string) []string {
  processed := make([]string, len(files))
  for i, v := range files {
    processed[i] = procRegexpReplace(v, options)
  }
  return processed
}

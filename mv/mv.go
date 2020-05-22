package main

import (
  "github.com/rowanho/go_cli/utilsCli"

  "os"
  "flag"
  "fmt"
  "regexp"
  "io/ioutil"
  "path/filepath"
)

/*
* Carries out the actual file replacement
*/

func directReplace(src string, dst string, prompt bool) (bool, error) {
  fileInfo, _:= os.Stat(dst)
  if fileInfo.IsDir() {
    dst = filepath.Join(dst, filepath.Base(src))
  }
  _, err := os.Stat(src)
  if os.IsNotExist(err) {
      return false, err
  }

  dstExist := false
  _, err = os.Stat(src)
  if !os.IsNotExist(err) {
      dstExist = true
  }

  if prompt && dstExist {
    overwrite := utilsCli.YesNoPrompt(fmt.Sprintf("%s already exists, overwrite? y/n: ", dst))
    if !overwrite {
      return false, nil
    }
  }

  err = os.Rename(src, dst)
  if err != nil {
    return false, err
  }
  return true, nil
}


/*
* Walks through all matching patterns, moves file if it matches
*/
func getFileSet(srcPatterns []string)  (map[string]bool, error) {
  replaceFiles := make(map[string]bool)
  err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
			return err
		}
    for _, pattern := range srcPatterns {
      isMatch, matchErr := regexp.MatchString(pattern, path)
      if matchErr == nil && isMatch {
        if !info.IsDir() {
            replaceFiles[path] = true
        } else {
            fns,_ := ioutil.ReadDir(path)
            for _, f := range fns {
              replaceFiles[f.Name()] = true
            }
        }
      }
    }
    return nil
  });
  return replaceFiles, err
}

func main() {
  iPtr := flag.Bool("i", false, "Prompts user when overwriting existing file")
  flag.Parse()
  args := flag.Args()
  if len(args) < 2 {
    fmt.Println("Error: please provide at least 2 arguments, source(s) and destination")
    return
  }
  srcs := args[:len(args) -1]
  dst := args[len(args) - 1]

  prompt := *iPtr
  fmt.Println(prompt)
  fileSet, fileErr := getFileSet(srcs)

  if fileErr != nil {
    fmt.Println(fileErr.Error())
    return
  }

  dstInfo, dstErr := os.Stat(dst)
  if dstErr != nil {
    fmt.Println(dstErr.Error())
    return
  } else if len(fileSet)> 1 && !dstInfo.IsDir() {
      fmt.Printf("%s needs to be a directory to move multiple files\n", dst)
      return
  }
  for filePath := range fileSet {
    written, err := directReplace(filePath, dst, prompt)
    if err != nil {
        fmt.Println(err.Error())
    } else if written {
        fmt.Printf("Moved %s to %s\n", filePath, dst)
    } else {
        fmt.Println("Cancelled overwrite")
    }
  }
}

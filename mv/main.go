package main

import (
  "github.com/rowanho/go_cli/utilsCli"

  "os"
  "flag"
  "fmt"
  "log"
)

/*
* Carries out the actual file replacement
*/

func directReplace(src string, dst string, prompt bool) (bool, error) {
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

func Execute() {
  src := os.Args[1]
  dst := os.Args[2]

  iPtr := flag.Bool("i", false, "Prompts user when overwriting existing file")
  flag.Parse()
  prompt := *iPtr
  fmt.Println(prompt)
  written, err := directReplace(src, dst, prompt)
  if err != nil {
    log.Fatal(err)
  } else if written {
    fmt.Println(fmt.Sprintf("Moved %s to %s", src, dst))
  } else {
    fmt.Println("Cancelled overwrite")
  }
}

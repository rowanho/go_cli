package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "flag"
)

var extraCaretOrder = []string{"^[", "^\\", "^]", "^^", "^_"}
var newline = 10
var tab = 9

/*
* Returns the caret notation of non-printable characters,
apart from tab or newline
*/
func caretNotation(b byte) string {
  i := int(b)
  if i < 27 {
      return "^" + string(byte(i + 64))
  } else if i < 32 {
      return extraCaretOrder[i-27]
  } else if i == 127 {
      return "^?"
  } else if i == 0 {
      return "^@"
  }
  return string(b)
}

func main() {
  bPtr := flag.Bool("b", false, "Number nonempty output lines.")
  EPtr := flag.Bool("E", false, "Display $ at the end of each line.")
  sPtr := flag.Bool("s", false, "Suppress repeated empty output lines.")
  TPtr := flag.Bool("T", false, "Display TAB characters as ^I.")
  vPtr := flag.Bool("v", false, "Displays nonprinting characters, except for tabs and the end of line character")
  flag.Parse()
  args := flag.Args()
  if len(args) == 0 {
    fmt.Println("Error: Need a filename argument")
    return
  }

  fileContents, readErr := ioutil.ReadFile(args[0])
  if readErr != nil {
    fmt.Println(readErr.Error())
    return
  }
  stringContents := string(fileContents)
  if *sPtr {
    lines := strings.Split(stringContents, "\n")
    newLines := make([]string, 0)
    subsq := 0
    for _, line := range lines {
      if strings.TrimSpace(line) == "" {
        subsq += 1
      } else {
        subsq = 0
      }
      if subsq < 2 {
        newLines = append(newLines, line)
      }
    }
    stringContents = strings.Join(newLines, "\n")
  }

  if !*EPtr && !*TPtr && !*vPtr && !*bPtr {
    fmt.Printf(stringContents)
    return
  }

  stringArr := make([]string, len(stringContents))
  for i, c := range stringContents {
    if int(byte(c)) == newline {
        if *EPtr {
          stringArr[i] += "$"
        }
        if *bPtr {
          stringArr[i] += ("  " + string(i + 1))
        }
        stringArr[i] += "\n"
    } else if *TPtr && int(byte(c)) == tab {
        stringArr[i] = "^I"
    } else if *vPtr  {
        stringArr[i] = caretNotation(byte(c))
    } else {
        stringArr[i] = string(c)
    }
  }
  fmt.Printf(strings.Join(stringArr, ""))
}

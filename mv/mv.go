package main

import (
	"flag"
	"fmt"
	"gocli/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// Carries out the actual file replacement

func directReplace(src string, dst string, prompt bool, prevent bool, update bool) (bool, error) {
	fileInfo, _ := os.Stat(dst)
	if fileInfo.IsDir() {
		dst = filepath.Join(dst, filepath.Base(src))
	}
	srcInfo, srcErr := os.Stat(src)
	if srcErr != nil {
		return false, srcErr
	}

	dstExist := false
	dstInfo, dstErr := os.Stat(dst)
	if dstErr == nil {
		dstExist = true
	}

	if dstExist {
		if prevent {
			return false, nil
		} else if update {
			if dstInfo.ModTime().After(srcInfo.ModTime()) {
				return false, nil
			}
		}

		if prompt {
			overwrite := utilsCli.YesNoPrompt(fmt.Sprintf("%s already exists, overwrite? y/n: ", dst))
			if !overwrite {
				return false, nil
			}
		}
	}

	err := os.Rename(src, dst)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Walks through all matching patterns, moves file if it matches
func getFileSet(srcPatterns []string) (map[string]bool, error) {
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
					fns, _ := ioutil.ReadDir(path)
					for _, f := range fns {
						replaceFiles[f.Name()] = true
					}
				}
			}
		}
		return nil
	})
	return replaceFiles, err
}

func main() {
	iPtr := flag.Bool("i", false, "Prompts user when overwriting existing file")
	nPtr := flag.Bool("n", false, "Prevents overwrite of existing file")
	uPtr := flag.Bool("u", false, "Updates when source is newer than destination")
	vPtr := flag.Bool("v", false, "Verbose - output destination of each file")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Error: please provide at least 2 arguments, source(s) and destination")
		return
	}
	srcs := args[:len(args)-1]
	dst := args[len(args)-1]

	fileSet, fileErr := getFileSet(srcs)

	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}

	dstInfo, dstErr := os.Stat(dst)
	if dstErr != nil {
		fmt.Println(dstErr.Error())
		return
	} else if len(fileSet) > 1 && !dstInfo.IsDir() {
		fmt.Printf("%s needs to be a directory to move multiple files\n", dst)
		return
	}
	for filePath := range fileSet {
		written, err := directReplace(filePath, dst, *iPtr, *nPtr, *uPtr)
		if err != nil {
			fmt.Println(err.Error())
		} else if written && *vPtr {
			fmt.Printf("Moved %s to %s\n", filePath, dst)
		} else if *vPtr {
			fmt.Printf("Cancelled move of %s to %s\n", filePath, dst)
		}
	}
}

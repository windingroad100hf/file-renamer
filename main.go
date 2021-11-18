package main

import (
	//"os"
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

type file string

func main() {
	var files []fs.FileInfo = GetFileNames()
	for _, fs := range files {
		if HasBadCharacter(fs.Name()) {
			if UserHasApproved(fs.Name()) {
				fmt.Println("Renaming...")
				RenameFile(fs.Name())
			}
		}
	}
}

func UserHasApproved(filename string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like to rename the following file: ", filename, "to ", getNewName(filename))
	userinput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return userinput == "yes\n" || userinput == "y\n"
}

func GetFileNames() []fs.FileInfo {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	return files
}

func HasBadCharacter(filename string) bool {
	return strings.Contains(filename, "_") || strings.Contains(filename, " ") || hasUpperCase(filename)

}

func hasUpperCase(filename string) bool {
	return filename != strings.ToLower(filename)
}

func getNewName(oldName string) string {
	newName := strings.ToLower(oldName)
	newName = strings.ReplaceAll(newName, "_", "-")
	newName = strings.ReplaceAll(newName, " ", "-")
	return newName

}
func RenameFile(oldName string) {
	newName := getNewName(oldName)
	os.Rename(oldName, newName)
}

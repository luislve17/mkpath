package utils

import (
	"fmt"
	"os"
	"strings"
)

func createFile(filePath string) {
	currentBuiltPath := ""
	if isRootedFolder := filePath[0] == '/'; isRootedFolder {
		fmt.Println("mkpath-DEBUG: Rooted folder")
		currentBuiltPath += "/"
	}
	fileFolderSections := strings.Split(filePath, "/")
	dirCreationStop := len(fileFolderSections) - 1
	isDirectory := filePath[len(filePath)-1] == '/'
	if isDirectory {
		dirCreationStop += 1
	}

	for _, folder := range fileFolderSections[:dirCreationStop] {
		currentBuiltPath += folder + "/"
		if !pathExists(currentBuiltPath) {
			if err := os.Mkdir(currentBuiltPath, os.ModePerm); err != nil {
				panic(err)
			}
		}
	}

	if !isDirectory {
		currentBuiltPath := currentBuiltPath + fileFolderSections[len(fileFolderSections)-1]
		if _, err := os.Create(currentBuiltPath); err != nil {
			panic(err)
		}
	}
	fmt.Printf("mkpath-MSG: Created: %s.\n", currentBuiltPath)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

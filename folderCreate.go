package main

import "os"

func createFolderIfNotExist(folderName string) {
	err := os.Mkdir(folderName, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

package utils

import (
	"bufio"
	"os"
)

type FileIo struct {
	filePath string
}

func (f *FileIo) Create() {
	_, err := os.Stat(f.filePath)
	if os.IsNotExist(err) {
		out, err := os.Create(f.filePath)
		if err != nil {
			panic(err)
		}
		defer out.Close()
	}
}
func (f *FileIo) Clear() {
	file, err := os.OpenFile(f.filePath, os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
func (f *FileIo) Write(data string) {
	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err2 := file.WriteString(data)

	if err2 != nil {
		panic(err2)
	}
}
func (f *FileIo) Read() []string {
	readFile, err := os.Open(f.filePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}

func NewFileIo(filePath string) *FileIo {
	return &FileIo{
		filePath: filePath,
	}
}

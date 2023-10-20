package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/f3rcho/gofromscratch/exercise"
)

var fileName string = "./files/txt/table.txt"

func SaveTable() {
	var text string = exercise.MultiplyTable()
	archive, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating archive" + err.Error())
		return
	}
	fmt.Fprintln(archive, text)
	archive.Close()
}

func CreateArchive() {
	var text string = exercise.MultiplyTable()
	if !Append(fileName, text) {
		fmt.Println("Error creating archive")
	}
}

func Append(fileName, text string) bool {
	arch, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening archive" + err.Error())
		return false
	}

	_, err = arch.WriteString(text + "\n")
	if err != nil {
		fmt.Println("Error writing archive" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func ReadArchive() {
	arch, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening archive" + err.Error())
		return
	}
	scanner := bufio.NewScanner(arch)

	for scanner.Scan() {
		registry := scanner.Text()
		fmt.Println("> " + registry + "\n")
	}
}

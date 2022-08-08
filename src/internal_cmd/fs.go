package internal_cmd

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile() string {
	path, err := os.Getwd()
	if err != nil {
		panic("Error! on ReadFile ex")
	}
	fullFileName := fmt.Sprintf("%s/%s", path, FILENAME)
	fmt.Println(fmt.Sprintf("[BUVETTE]: BuvetteFile PATH: %s", fullFileName))
	file, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		fmt.Println("[BUVETTE]: Sorry! No BuvetteFile file here :(", path)
		return ""
	}
	return string(file)
}

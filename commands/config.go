package commands

import (
	"os"
	"fmt"
)

const (
	BUFSIZE = 1024
	GIT_FILE_PATH = "./.git"
)

func PrintConfig() {
	file, err := os.Open(GIT_FILE_PATH + "/config")
	if err != nil {
		fmt.Println("[ERROR]: config file doesn't exists.")
	} else {
		buf := make([]byte, BUFSIZE)
		for {
			n, err := file.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Println("[ERROR]: An error occured while reading confg file.")
			}

			fmt.Print(string(buf[:n]))
		}

	}

}
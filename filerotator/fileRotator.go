package filerotator

import (
	"fmt"
	"os"
	"time"
)

const MAX_LINES_PER_FILE  = 1000
const LOG_LINE = "The quick brown fox jumps over the lazy dog"
const LOG_FILE_NAME = "SampleFile"
const FILE_EXTENSION = ".txt"
const FILE_DIRECTORY = "generatedfiles/"
const LOG_FILE_PATH = FILE_DIRECTORY+LOG_FILE_NAME+FILE_EXTENSION
var rotatedFileCounter int = 0


func GenerateAndRotateFile() {

	file,err := os.Create(LOG_FILE_PATH)
	if err != nil {
		fmt.Println("Cannot Create file!"+err.Error())
	}
	defer file.Close()

	totalLinesWritten := 0

	for {

		file.Write([]byte(LOG_LINE+"\n"))
		totalLinesWritten++

		if totalLinesWritten == MAX_LINES_PER_FILE {
			file.Close()
			err := os.Rename(LOG_FILE_PATH, getRotatedFilePath())
			if err != nil {
				fmt.Println("Error: Cannot Rename file!"+err.Error())
			}
			file,err = os.Create(LOG_FILE_PATH)
			totalLinesWritten = 0
		}
		time.Sleep(300)

	}


}

func getRotatedFilePath() string {

	rotatedFileCounter++
	return fmt.Sprintf(FILE_DIRECTORY+LOG_FILE_NAME+"%v"+FILE_EXTENSION,rotatedFileCounter)

}


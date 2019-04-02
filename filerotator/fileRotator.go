package filerotator

import (
	"fmt"
	"os"
)

const MAX_LINES_PER_FILE  = 100
const LOG_LINE = "The quick brown fox jumps over the lazy dog"


func GenerateAndRotateFile() {

	file,err := os.Create("generatedfiles/SampleFile.txt")
	if err != nil {
		fmt.Println("Cannot Create file!"+err.Error())
	}
	defer file.Close()

	totalLinesWritten := 0

	for {

		file.Write([]byte(LOG_LINE+"\n"))
		totalLinesWritten++

		if totalLinesWritten == MAX_LINES_PER_FILE {
			break
		}

	}






}


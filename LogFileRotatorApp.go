package main

import (
	"fmt"
	"go-logfile-rotator/filerotator"
	"time"
)

func main() {

	fmt.Println("Starting File Rotator...")
	time.Sleep(1000)
	filerotator.GenerateAndRotateFile()

}

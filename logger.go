package logger

import "fmt"

var Version = "1.0"

func Log(message string) {
	fmt.Println("[log]" + message)
}

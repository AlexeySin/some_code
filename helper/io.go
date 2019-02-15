package helper

import (
	"fmt"
	"log"
	"os"
)

func Out(data ...interface{}) {
	os.Stdout.Write([]byte(fmt.Sprintf("%v\n", data)))
}

func Warning(message string) {
	log.Printf("WARNING: %s\n\n", message)
}

package writer

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func WriteTo(file *os.File, data string) {
	numBytes, err := file.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bytes written: ", numBytes)
}

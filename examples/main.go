package main

import (
	"bufio"
	"log"
	"os"

	utils "github.com/Safwanseban/go-utils"
)

func main() {
	//initializing a reader for sample utilzation
	rw := utils.RW{}
	r := bufio.NewReader(os.Stdin)
	readedData := rw.WithReader(r).Read()

	//initializing the writer for writer
	writer := rw.WithWriter(bufio.NewWriter(os.Stdout))
	len, err := writer.Write(readedData, 5)
	if err != nil {
		log.Println("error writing buffer")
	}
	log.Println("availble buffer to write ", len)

}

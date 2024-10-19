package main

import (
	"chat-app/db"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not connect to Database :%s", err)
	}

	name := flag.String("name", "world", "The name to greet.")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("Hello, %s!\n", *name)
	} else if flag.Arg(0) == "list" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Printf("Hello, %s!\n", *name)
	}
}

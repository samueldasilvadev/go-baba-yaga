package lua

import (
	"fmt"
	"log"
	"os"
)

func LoadScripts() {
	files, err := os.ReadFile("./lua/addJob.lua")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(files))
}

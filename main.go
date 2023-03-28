package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Queue struct {
	Redis   *redis.Client
	Context context.Context
}

func LoadScripts() {
	files, err := os.ReadFile("./lua/addJob.lua")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(files))
}

func (q *Queue) addJob() {}

func main() {
	LoadScripts()
	q := Queue{}
	q.addJob()
}

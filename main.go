package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/samueldasilvadev/go-baba-yaga/lua"
)

type Queue struct {
	Redis   *redis.Client
	Context context.Context
}

func (q *Queue) addJob() {}

func main() {
	lua.LoadScripts()
	q := Queue{}
	q.addJob()
}

package redis

import (
	"context"
	"os"

	"github.com/mediocregopher/radix/v4"
)

var Client radix.Client

func init() {
	connectionAddr := os.Getenv("REDIS_CONNECTION")

	Client, _ = (radix.PoolConfig{}).New(context.TODO(), "tcp", connectionAddr)
}

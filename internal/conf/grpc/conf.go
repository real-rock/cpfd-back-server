package grpc

import (
	"context"
	"fmt"
	"os"
)

type conf struct {
	host string
	port string
	ctx  context.Context
}

func newConf() *conf {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_INSECURE_PORT")

	return &conf{
		host: host,
		port: port,
		ctx:  context.Background(),
	}
}

func (c *conf) getDSN() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

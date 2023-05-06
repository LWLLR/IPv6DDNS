package main

import (
	"context"
	"flag"
)

var ConfigPath string

func init() {
	flag.StringVar(&ConfigPath, "c", "./config.yaml", "配置地址")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	t, err := NewTencentCloudUpdater()
	if err != nil {
		panic(err)
	}
	if err = Watch(ctx, t); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
)

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

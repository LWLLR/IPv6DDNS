package main

import (
	"context"
	"fmt"
	"time"
)

type IPV6DDNS interface {
	Update(ctx context.Context, ipv6 string) error
}

func Watch(ctx context.Context, updater IPV6DDNS) error {
	if updater == nil {
		return fmt.Errorf("updater is nil")
	}
	t := time.NewTicker(time.Second * time.Duration(Conf.Interval))
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			updater.Update(ctx, GetIPV6())
		}
	}
	return nil
}

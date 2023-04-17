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
	t := time.NewTicker(time.Second * time.Duration(Conf.GetInterval()))
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			ip, err := GetIPV6()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err = updater.Update(ctx, ip); err != nil {
				fmt.Println("update dns err:", err)
			}
		}
	}
}

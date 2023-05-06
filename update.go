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
	if err := update(ctx, updater); err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			if err := update(ctx, updater); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func update(ctx context.Context, updater IPV6DDNS) error {
	ip, err := GetIPV6()
	if err != nil {
		return err
	}
	if err = updater.Update(ctx, ip); err != nil {
		fmt.Println("update dns err:", err)
	}
	return nil
}

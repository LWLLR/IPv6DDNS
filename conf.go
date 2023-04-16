package main

var Conf Config

type Config struct {
	Interval int
}

func init() {
	Conf = Config{
		Interval: 3,
	}
}

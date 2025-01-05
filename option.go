package airouteros

import "time"

type Options struct {
	Address       string
	Username      string
	Password      string
	Timeout       time.Duration
	AutoReconnect bool
}

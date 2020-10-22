package socks5

import "time"

type scontext struct {
	AuthContext *AuthContext
}

func (ctx *scontext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ctx *scontext) Done() <-chan struct{} {
	return nil
}

func (ctx *scontext) Err() error {
	return nil
}

func (ctx *scontext) Value(key interface{}) interface{} {
	return nil
}

func (ctx *scontext) String() string {
	return "socks5 context"
}

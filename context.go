package socks5

import "time"

//Context socks5 context
type Context struct {
	AuthContext *AuthContext
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ctx *Context) Done() <-chan struct{} {
	return nil
}

func (ctx *Context) Err() error {
	return nil
}

func (ctx *Context) Value(key interface{}) interface{} {

	return nil
}

func (ctx *Context) String() string {
	return "socks5 context"
}

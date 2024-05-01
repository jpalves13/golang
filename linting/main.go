package main

import "net/http/httputil"

func main() {
	lintByPass()
	lintFail()
}

func lintFail() {
	httputil.NewClientConn(nil, nil)
}

func lintByPass() {
	httputil.NewClientConn(nil, nil) //nolint
}

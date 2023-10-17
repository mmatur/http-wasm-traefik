package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = handleRequest
}

// handleRequest implements a simple HTTP router.
func handleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	req.Headers().Add("X-POWPOW", "hello")
	// proceed to the next handler on the host.
	next = true
	return
}

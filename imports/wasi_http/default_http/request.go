package default_http

import (
	"context"
	"log"

	"github.com/stealthrocket/wasi-go/imports/wasi_http/common"
	"github.com/stealthrocket/wasi-go/imports/wasi_http/types"
	"github.com/tetratelabs/wazero/api"
)

type Handler struct {
	req      *types.Requests
	res      *types.Responses
	fields   *types.FieldsCollection
	settings *common.Settings
}

// Request handles HTTP serving. It's currently unimplemented
func requestFn(_ context.Context, mod api.Module, a, b, c, d, e, f, g, h, j, k, l, m, n, o uint32) int32 {
	return 0
}

// Handle handles HTTP client calls.
// The remaining parameters (b..h) are for the HTTP Options, currently unimplemented.
func (handler *Handler) handleFn(_ context.Context, mod api.Module, request, b, c, d, e, f, g, h uint32) uint32 {
	req, ok := handler.req.GetRequest(request)
	if !ok {
		log.Printf("Failed to get request: %v\n", request)
		return 0
	}
	if !handler.settings.IsMethodAllowed(req.Method) {
		log.Printf("method not allowed: (%s)\n", req.Method)
		return 0
	}
	if !handler.settings.IsHostAllowed(req.Authority) {
		log.Printf("host not allowed: (%s)\n", req.Authority)
		return 0
	}
	r, err := req.MakeRequest(handler.fields)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return handler.res.MakeResponse(r)
}

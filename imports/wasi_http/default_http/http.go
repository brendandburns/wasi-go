package default_http

import (
	"context"

	"github.com/stealthrocket/wasi-go/imports/wasi_http/common"
	"github.com/stealthrocket/wasi-go/imports/wasi_http/types"
	"github.com/tetratelabs/wazero"
)

const ModuleName = "default-outgoing-HTTP"

func Instantiate(ctx context.Context, r wazero.Runtime, req *types.Requests, res *types.Responses, f *types.FieldsCollection, s *common.Settings) error {
	handler := &Handler{req, res, f, s}
	_, err := r.NewHostModuleBuilder(ModuleName).
		NewFunctionBuilder().WithFunc(requestFn).Export("request").
		NewFunctionBuilder().WithFunc(handler.handleFn).Export("handle").
		Instantiate(ctx)
	return err
}

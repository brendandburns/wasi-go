package types

import (
	"context"
	"fmt"

	"github.com/stealthrocket/wasi-go/imports/wasi_http/common"
	"github.com/stealthrocket/wasi-go/imports/wasi_http/streams"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

const ModuleName = "wasi:http/types@0.2.0-rc-2023-10-18"

func logFn(ctx context.Context, mod api.Module, ptr, len uint32) {
	str, _ := common.ReadString(mod, ptr, len)
	fmt.Print(str)
}

func Instantiate(ctx context.Context, rt wazero.Runtime, s *streams.Streams, r *Requests, rs *Responses, f *FieldsCollection, o *OutResponses) error {
	_, err := rt.NewHostModuleBuilder(ModuleName).
		NewFunctionBuilder().WithFunc(r.newOutgoingRequestFn).Export("new-outgoing-request").
		NewFunctionBuilder().WithFunc(f.newFieldsFn).Export("[constructor]fields").
		NewFunctionBuilder().WithFunc(f.dropFieldsFn).Export("drop-fields").
		NewFunctionBuilder().WithFunc(f.fieldsEntriesFn).Export("fields-entries").
		NewFunctionBuilder().WithFunc(r.dropOutgoingRequestFn).Export("drop-outgoing-request").
		NewFunctionBuilder().WithFunc(r.outgoingRequestWriteFn).Export("outgoing-request-write").
		NewFunctionBuilder().WithFunc(rs.dropIncomingResponseFn).Export("drop-incoming-response").
		NewFunctionBuilder().WithFunc(rs.incomingResponseStatusFn).Export("incoming-response-status").
		NewFunctionBuilder().WithFunc(rs.incomingResponseHeadersFn).Export("incoming-response-headers").
		NewFunctionBuilder().WithFunc(rs.incomingResponseConsumeFn).Export("incoming-response-consume").
		NewFunctionBuilder().WithFunc(futureResponseGetFn).Export("future-incoming-response-get").
		NewFunctionBuilder().WithFunc(r.incomingRequestMethodFn).Export("[method]incoming-request.method").
		NewFunctionBuilder().WithFunc(r.incomingRequestPathFn).Export("[method]incoming-request.path-with-query").
		NewFunctionBuilder().WithFunc(r.incomingRequestAuthorityFn).Export("[method]incoming-request.authority").
		NewFunctionBuilder().WithFunc(r.incomingRequestHeadersFn).Export("[method]incoming-request.headers").
		NewFunctionBuilder().WithFunc(incomingRequestConsumeFn).Export("[method]incoming-request.consume").
		NewFunctionBuilder().WithFunc(r.dropIncomingRequestFn).Export("drop-incoming-request").
		NewFunctionBuilder().WithFunc(o.setResponseOutparamFn).Export("[static]response-outparam.set").
		NewFunctionBuilder().WithFunc(rs.newOutgoingResponseFn).Export("[constructor]outgoing-response").
		NewFunctionBuilder().WithFunc(rs.outgoingResponseWriteFn).Export("[method]outgoing-response.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyWriteFn).Export("[method]outgoing-body.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyFinishFn).Export("[static]outgoing-body.finish").
		NewFunctionBuilder().WithFunc(dropOutgoingResponseFn).Export("drop-outgoing-response").
		NewFunctionBuilder().WithFunc(logFn).Export("log-it").
		Instantiate(ctx)
	return err
}

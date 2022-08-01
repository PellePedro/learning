
# Processor

Implemented Code
```
handler := NewCartServiceHandler()
processor := generated.NewCartServiceProcessor(handler)
server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

```

Generated Code
```

func NewCartServiceProcessor(handler CartService) *CartServiceProcessor {
  self15 := &CartServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self15.processorMap["AddItem"] = &cartServiceProcessorAddItem{handler:handler}
  self15.processorMap["GetCart"] = &cartServiceProcessorGetCart{handler:handler}
  self15.processorMap["EmptyCart"] = &cartServiceProcessorEmptyCart{handler:handler}
   return self15
}

func (p *CartServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x16 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x16.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x16
}

```

# GRPC
```

func (p *MockProto) BuildMockService(server *MockGrpcServer, endpoint *types.Endpoint) error {
  ...
  serviceDesc := p.fileDescriptor.FindService(fmt.Sprintf("%s.%s", p.ProtoPackageName, serviceName))
	if len(methods) != 0 {
		grpcServiceDesc.Methods = methods
	}
	if len(streams) != 0 {
		grpcServiceDesc.Streams = streams
	}
  ...
	log.Infof("Building service %s is done", serviceDesc.GetName())
	server.Server.RegisterService(grpcServiceDesc, struct{}{})

}
```

```
type ServiceDesc struct {
	ServiceName string
	// The pointer to the service interface. Used to check whether the user
	// provided implementation satisfies the interface requirements.
	HandlerType interface{}
	Methods     []MethodDesc
	Streams     []StreamDesc
	Metadata    interface{}
}
``

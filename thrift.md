
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
```

# Thrift
## Processor

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





### Phython Client

```
from thrift import Thrift
from thrift.transport import TSSLSocket
from thrift.transport import TSocket
from thrift.transport import THttpClient
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

from demo import ProductCatalogService

def main():
    uri = '{0}://{1}:{2}{3}'.format('http','product-catalog-service', 50000, '/ProductCatalogService')

   /* HTTP Client */
    transport = THttpClient.THttpClient(uri)
    protocol = TBinaryProtocol.TBinaryProtocolFactory().getProtocol(transport)

	# TLS
    # socket = TSSLSocket.TSSLSocket('product-catalog-service', 50000, validate=False)
    # transport = TTransport.TBufferedTransport(socket)
    # protocol = TBinaryProtocol.TBinaryProtocol(transport)

   transport.open()
    print("successfully connected to product-catalog-server", flush=True)
    client = ProductCatalogService.Client(protocol)
    products = client.ListProducts()
    print(f"successfully retrieved products from product cataloge {products}", flush=True)
    transport.close()


from thrift.transport import THttpClient
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

from tutorial import Calculator

transport = THttpClient.THttpClient('https://your-service.com')
transport = TTransport.TBufferedTransport(transport)
protocol = TBinaryProtocol.TBinaryProtocol(transport)
client = Calculator.Client(protocol)

# Connect!
transport.open()
client.ping()
```


```
import os
import sys
sys.path.append('thriftpy')

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

from demo import RecommendationService

def main():

    target = 'recommendation-service'
    transport = TSocket.TSocket(target, 50000)
    transport = TTransport.TBufferedTransport(transport)
    protocol = TBinaryProtocol.TBinaryProtocol(transport)
    transport.open()

    client = RecommendationService.Client(protocol)
    # Call ListRecommendations with one product in cart
    res = client.ListRecommendations("6E92ZMYYFZ")
    print(res)
    transport.close()

  transport = TSSLSocket.TSSLSocket(self.productCatalogServerHost, self.productCatalogServerPort, cert_reqs=ssl.CERT_NONE)

if __name__ == '__main__':
    try:
        main()
    except Thrift.TException as tx:
        print('%s' % tx.message)
```


### Go Client
```
	opt := NewDefaultOption()
	opt.Secure = true
	opt.Buffered = true
	opt.HttpTransport = true
	opt.HttpUrl = "/ProductCatalogService"
	hostPort := "127.0.0.1:50000"
	client, trans, err := NewThriftClient(hostPort, opt)
	assert.Nil(t, err)
	c := demo.NewProductCatalogServiceClient(client)
	err = trans.Open()
	assert.Nil(t, err)
	allproducts, err := c.ListProducts(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 9, len(allproducts))
	jsonProducts, err := json.Marshal(allproducts)
	ioutil.WriteFile("products.txt", []byte(jsonProducts), 0644)
	_ = jsonProducts
	products, err := c.SearchProducts(context.TODO(), "kitchen")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
	product, err := c.GetProduct(context.TODO(), product_id)
	assert.Nil(t, err)
	_ = product


```
### Thrift Stack
```
/*
 * Copyright Skyramp Authors 2022
 */
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/apache/thrift/lib/go/thrift"
)

type ThriftProtocolType string

const (
	Binary     ThriftProtocolType = "binary"
	Json       ThriftProtocolType = "json"
	SimpleJson ThriftProtocolType = "simplejson"
	Compact    ThriftProtocolType = "compact"
)

type Option struct {
	HttpTransport bool
	HttpUrl       string
	Protocol      ThriftProtocolType
	Secure        bool
	Buffered      bool
	Framed        bool
}

func NewDefaultOption() *Option {
	return &Option{
		Protocol:      Binary,
		HttpTransport: false,
		Secure:        true,
		Buffered:      true,
		Framed:        false,
	}
}

var (
	protocolFactoryMap          = make(map[ThriftProtocolType]thrift.TProtocolFactory)
	bufferedTransportFactoryMap = make(map[bool]thrift.TTransportFactory)
)

func init() {
	protocolFactoryMap[Binary] = thrift.NewTBinaryProtocolFactoryConf(nil)
	protocolFactoryMap[Json] = thrift.NewTJSONProtocolFactory()
	protocolFactoryMap[SimpleJson] = thrift.NewTSimpleJSONProtocolFactoryConf(nil)
	protocolFactoryMap[Compact] = thrift.NewTCompactProtocolFactoryConf(nil)
	bufferedTransportFactoryMap[true] = thrift.NewTBufferedTransportFactory(8192)
	bufferedTransportFactoryMap[false] = thrift.NewTTransportFactory()
}

/*
 * Creats a ThriftServer with autogenerated certificate
 */
func NewHttpThriftServer(addr string, opt *Option, processor thrift.TProcessor) error {
	protocolFactory := protocolFactoryMap[opt.Protocol]
	if !opt.HttpTransport {
		return fmt.Errorf("NewHttpThriftServer called with opt.httpTransport set to [false]")
	}

	http.HandleFunc(opt.HttpUrl, thrift.NewThriftHandlerFunc(processor, protocolFactory, protocolFactory))
	log.Infof("tcp listener opened for thrift socket %s", addr)
	var err error
	if opt.Secure {
		err = startHttps(addr)
	} else {
		err = http.ListenAndServe(addr, nil)
	}
	if err != nil {
		log.Errorf("failed to start http server: %v", err)
	}
	return nil
}

/*
 * Creats a ThriftServer with provided key and certificate
 */
func NewHttpThriftServer2(addr string, opt *Option, certFile, keyFile string, processor thrift.TProcessor) error {
	protocolFactory := protocolFactoryMap[opt.Protocol]
	if !opt.HttpTransport {
		return fmt.Errorf("NewHttpThriftServer called with opt.httpTransport set to [false]")
	}
	http.HandleFunc(opt.HttpUrl, thrift.NewThriftHandlerFunc(processor, protocolFactory, protocolFactory))
	var err error
	if opt.Secure {
		err = http.ListenAndServeTLS(addr, certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(addr, nil)
	}
	if err != nil {
		fmt.Printf("Failed to start http server: %v\n", err)
	}
	return nil
}

func NewStandardThriftServer(addr string, opt *Option, processor thrift.TProcessor) error {
	log.Infof("creating a mocked thrift service on %s with options:", addr)
	log.Infof("protocol: [%s], httpTransport: [%t], secure: [%t], buffered: [%t], framed: [%t]",
		opt.Protocol, opt.HttpTransport, opt.Secure, opt.Buffered, opt.Framed)
	protocolFactory, ok := protocolFactoryMap[opt.Protocol]
	if !ok {
		return fmt.Errorf("no protocol found for NewThriftServer %s", opt.Protocol)
	}

	var transportFactory thrift.TTransportFactory
	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	transportFactory = bufferedTransportFactoryMap[opt.Buffered]

	if opt.Framed {
		transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, cfg)
	}
	var transport thrift.TServerTransport
	var err error
	if opt.Secure {
		serverTLSConf, clientTLSConf, caPEM, err := generateCertificate()
		_, _ = clientTLSConf, caPEM
		if err != nil {
			return fmt.Errorf("failed to create tls certificate %w", err)
		}
		transport, err = thrift.NewTSSLServerSocket(addr, serverTLSConf)
		if err != nil {
			return fmt.Errorf("failed to create tls certificate %w", err)
		}
	} else {
		transport, err = thrift.NewTServerSocket(addr)
		if err != nil {
			return fmt.Errorf("failed to create thrift server %w", err)
		}
	}
	log.Infof("tcp listener opened for thrift socket %s", addr)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()
	if err != nil {
		return fmt.Errorf("failed to start Thrift server... on port %s", addr)
	}
	log.Infof("thrift service... on port %s terminated \n", addr)
	return nil
}

// Starts https with in memory generated certificate
func startHttps(addr string) error {
	serverTLSConf, clientTLSConf, caPEM, _ := generateCertificate()
	_, _ = clientTLSConf, caPEM
	getCertificate := func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
		return &serverTLSConf.Certificates[0], nil
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: http.DefaultServeMux,
		TLSConfig: &tls.Config{
			MinVersion:     tls.VersionTLS13,
			GetCertificate: getCertificate,
		},
	}
	err := srv.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("failed to perform https listen and serve tls %w", err)
		return fmt.Errorf("failed to perform https listen and serve tls %w", err)
	}
	return nil
}

/*
*  Return a new ThriftClient
 */
func NewThriftClient(hostPort string, opt *Option) (client *thrift.TStandardClient, trans thrift.TTransport, err error) {

	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	protocolFactory := protocolFactoryMap[opt.Protocol]
	if opt.Secure {
		trans = thrift.NewTSSLSocketConf(hostPort, cfg)
	} else {
		trans = thrift.NewTSocketConf(hostPort, nil)
	}
	if err != nil {
		return nil, nil, err
	}
	if opt.HttpTransport {
		if opt.Secure {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			trans, err = thrift.NewTHttpClientWithOptions(fmt.Sprintf("https://%s%s", hostPort, opt.HttpUrl), thrift.THttpClientOptions{Client: client})
		} else {
			trans, err = thrift.NewTHttpClient(fmt.Sprintf("http://%s%s", hostPort, opt.HttpUrl))
		}
	} else {
		if opt.Buffered {
			trans = thrift.NewTBufferedTransport(trans, 8192)
		} else {
			trans = thrift.NewTFramedTransportConf(trans, cfg)
		}
	}
	if err != nil {
		return nil, nil, err
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client = thrift.NewTStandardClient(iprot, oprot)
	return
}

```

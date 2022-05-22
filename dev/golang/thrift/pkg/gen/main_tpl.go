package gen

const (
	ThriftProtocolBinary     = "binary"
	ThriftProtocolJson       = "json"
	ThriftProtocolSimpleJson = "simplejson"
	ThriftProtocolCompact    = "compact"
)

type ServerTplData struct {
	Address       string
	Protocol      string
	Framed        bool
	Buffered      bool
	Secure        bool
	MockedService string
}

/*
import (
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)
*/

var tplServer = `{{define "tplServer"}}
func main() {
   var protocolFactory thrift.TProtocolFactory
   {{- if (eq .Protocol "binary")     }}
   protocolFactory = thrift.NewTCompactProtocolFactoryConf(nil)
   {{- else if (eq .Protocol "json")       }}
   protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(nil)
   {{- else if (eq .Protocol "simplejson") }}
   protocolFactory = thrift.NewTJSONProtocolFactory()
   {{- else if (eq .Protocol "compact")    }}
   protocolFactory = thrift.NewTBinaryProtocolFactoryConf(nil)
   {{- end }}
   var transportFactory thrift.TTransportFactory
   cfg := &thrift.TConfiguration{
     TLSConfig: &tls.Config{
   	   InsecureSkipVerify: true,
     },
   }

   {{- if (eq .Buffered true ) }}
   transportFactory = thrift.NewTBufferedTransportFactory(8192)
   {{- else }}
     transportFactory = thrift.NewTTransportFactory()
   {{- end }}
   {{- if (eq .Framed true ) }}
     transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, cfg)
   {{- end }}

   var transport thrift.TServerTransport
   var err error
   {{- if .Secure }}
   cfg := new(tls.Config)
   if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
   	 cfg.Certificates = append(cfg.Certificates, cert)
   } else {
   	 return err
   }
   transport, err = thrift.NewTSSLServerSocket(addr, cfg)
   {{- else }}
   transport, err = thrift.NewTServerSocket( {{ .Address }})
   {{- end }}
   if err != nil {
     return err
   }

   server := thrift.NewTSimpleServer4(GetProcessor(), transport, transportFactory, protocolFactory)

   fmt.Println("Starting Thrift server... on ", addr)
   return server.Serve()
)
{{ end }}
`



# Python
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




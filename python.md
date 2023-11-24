## Remote development with VScode
[Remote Development](https://code.visualstudio.com/docs/remote/remote-overview)

## Create public ssh keypair on local machine
```
ssh-keygen -t rsa -b 4096
```
## Connect to remote from VScode

## Setting up a python project

### Create directory for project
```
mkdir projectA
cd projectA
```

### Create python virtual environment
[virtual environment](https://www.freecodecamp.org/news/how-to-setup-virtual-environments-in-python/)
```
python3 -m venv venv
```

### Activate the python virtual environment
```
. ./venv/bin/activate
```

### Sample
To create an HTTP server in Python and demonstrate how to parse JSON data, you can use the built-in `http.server` module for the server part and the `json` module for parsing JSON data. Here's a step-by-step guide along with an example:

Step 1: Import the required modules.
```python
import http.server
import socketserver
import json
```

Step 2: Define a request handler class that inherits from `http.server.BaseHTTPRequestHandler`. You'll override the `do_POST` method to handle POST requests and parse JSON data from the request body.

```python
class MyHandler(http.server.BaseHTTPRequestHandler):
    def do_POST(self):
        # Check if the request path is "/parse_json"
        if self.path == '/parse_json':
            # Get the content length from the request headers
            content_length = int(self.headers['Content-Length'])
            
            # Read the request body based on the content length
            request_body = self.rfile.read(content_length)
            
            try:
                # Parse the JSON data from the request body
                json_data = json.loads(request_body.decode('utf-8'))
                
                # Process the JSON data (you can customize this part)
                response_data = {'message': 'JSON data received and processed successfully'}
                response_json = json.dumps(response_data)
                
                # Send a response with a 200 OK status code
                self.send_response(200)
                self.send_header('Content-type', 'application/json')
                self.end_headers()
                self.wfile.write(response_json.encode('utf-8'))
                
            except json.JSONDecodeError:
                # If JSON parsing fails, send a 400 Bad Request response
                self.send_response(400)
                self.send_header('Content-type', 'text/plain')
                self.end_headers()
                self.wfile.write(b'Invalid JSON data')
        else:
            # If the request path is not "/parse_json", send a 404 Not Found response
            self.send_response(404)
            self.send_header('Content-type', 'text/plain')
            self.end_headers()
            self.wfile.write(b'Not Found')
```

Step 3: Create a server instance and specify the request handler class.

```python
PORT = 8000

with socketserver.TCPServer(("", PORT), MyHandler) as httpd:
    print("Server started at port", PORT)
    
    # Start the server and keep it running until interrupted
    httpd.serve_forever()
```

Now, you have a basic HTTP server that listens on port 8000 and handles POST requests to the "/parse_json" path. It parses JSON data from the request body and sends a response.

To test the server, you can use a tool like `curl` or create a simple Python client. Here's an example using `curl`:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"key": "value"}' http://localhost:8000/parse_json
```

This sends a POST request with JSON data to the server, and you should receive a response indicating the success or failure of JSON parsing.


[An overview of HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview)

HTTP is a protocol for fetching resources such as HTML documents. 
It is the foundation of any data exchange on the Web and it is a client-server protocol, 
which means requests are initiated by the recipient, usually the Web browser. 
A complete document is reconstructed from the different sub-documents fetched, 
for instance, text, layout description, images, videos, scripts, and more.


Clients and servers communicate by exchanging individual messages (as opposed to a stream of data). 
The messages sent by the client, usually a Web browser, are called requests and the messages sent 
by the server as an answer are called responses.

Designed in the early 1990s, HTTP is an extensible protocol which has evolved over time. 
It is an application layer protocol that is sent over TCP, or over a TLS-encrypted TCP connection, 
though any reliable transport protocol could theoretically be used. Due to its extensibility, 
it is used to not only fetch hypertext documents, but also images and videos or to post content 
to servers, like with HTML form results. HTTP can also be used to fetch parts of documents to update 
Web pages on demand.

## GPT
"Students, today we are discussing the HyperText Transfer Protocol or HTTP. This protocol is the backbone of all data exchange on the World Wide Web. It is a client-server type protocol where the client, typically a web browser, initiates the request for resources such as HTML documents. The complete document is constructed from various sub-documents like text, images, videos, scripts, and more.

HTTP operates by exchanging individual messages, not a stream of data, between the client and server. The client sends a request and the server responds with a response.

HTTP was developed in the early 1990s and has evolved over time. It is an application layer protocol that runs over a reliable transport protocol, such as TCP or a TLS-encrypted TCP connection. Additionally, due to its extensibility, it can be used for various purposes beyond just fetching HTML documents, such as posting content to servers, and updating web pages on demand. That's a brief overview of HTTP, any questions?"

## Components of HTTP-based systems
HTTP is a client-server protocol: requests are sent by one entity, the user-agent (or a proxy on behalf of it). 
Most of the time the user-agent is a Web browser, but it can be anything, for example, a robot that crawls the 
Web to populate and maintain a search engine index.

Each individual request is sent to a server, which handles it and provides an answer called the response. 
Between the client and the server there are numerous entities, collectively called proxies, 
which perform different operations and act as gateways or caches, for example.

In reality, there are more computers between a browser and the server handling the request: there are routers, 
modems, and more. Thanks to the layered design of the Web, these are hidden in the network and transport layers. 
HTTP is on top, at the application layer. Although important for diagnosing network problems, the underlying 
layers are mostly irrelevant to the description of HTTP.


## GPT
Good morning, students! I hope you are all ready for today's lecture on the HTTP protocol. HTTP, as you know, stands for HyperText Transfer Protocol and it is a fundamental component of the World Wide Web.

The protocol is based on the client-server architecture where the client, or user-agent, sends requests to the server, which then provides the corresponding response. The user-agent can be a web browser, a robot, or any other entity capable of sending requests.

In between the client and the server, there may be one or more intermediaries, commonly known as proxies, which perform various functions such as acting as gateways or caching content. These proxies can help improve the overall efficiency of the system by reducing the number of requests sent to the server and by storing frequently requested data.

It is also worth mentioning that there are many other components involved in the communication process, such as routers and modems, but they are hidden in the lower layers of the network and do not directly interact with the HTTP protocol. HTTP operates at the application layer and is essential for diagnosing network issues.

In conclusion, HTTP is a widely used protocol for communication and data transfer on the Web. It is a crucial aspect of the World Wide Web and an important topic for anyone looking to gain a deeper understanding of the inner workings of the Internet.



# Basic aspects of HTTP

## HTTP is simple
HTTP is generally designed to be simple and human-readable. HTTP messages can be read and understood by humans, 
providing easier testing for developers, and reduced complexity for newcomers.

## HTTP is extensible
Introduced in HTTP/1.0, HTTP headers make this protocol easy to extend and experiment with. New functionality 
can even be introduced by a simple agreement between a client and a server about a new header's semantics.

## HTTP is stateless, but not sessionless
HTTP is stateless: there is no link between two requests being successively carried out on the same connection. 
This immediately has the prospect of being problematic for users attempting to interact with certain pages coherently, 
for example, using e-commerce shopping baskets. But while the core of HTTP itself is stateless, HTTP cookies allow 
the use of stateful sessions. Using header extensibility, HTTP Cookies are added to the workflow, 
allowing session creation on each HTTP request to share the same context, or the same state.

## HTTP and connections
A connection is controlled at the transport layer, and therefore fundamentally out of scope for HTTP. 
HTTP doesn't require the underlying transport protocol to be connection-based; it only requires it to be reliable, 
or not lose messages (at minimum, presenting an error in such cases). 

Among the two most common transport protocols on the Internet, TCP is reliable and UDP isn't. HTTP therefore 
relies on the TCP standard, which is connection-based.

Before a client and server can exchange an HTTP request/response pair, they must establish a TCP connection, 
a process which requires several round-trips. The default behavior of HTTP/1.0 is to open a separate TCP connection 
for each HTTP request/response pair. 

This is less efficient than sharing a single TCP connection when multiple requests are sent in close succession.

In order to mitigate this flaw, HTTP/1.1 introduced pipelining (which proved difficult to implement) and persistent 
connections: the underlying TCP connection can be partially controlled using the Connection header. HTTP/2 went a step 
further by multiplexing messages over a single connection, helping keep the connection warm and more efficient.

Experiments are in progress to design a better transport protocol more suited to HTTP. For example, Google is 
experimenting with QUIC which builds on UDP to provide a more reliable and efficient transport protocol.

## What can be controlled by HTTP
This extensible nature of HTTP has, over time, allowed for more control and functionality of the Web. 
Cache and authentication methods were functions handled early in HTTP history. 
The ability to relax the origin constraint, by contrast, was only added in the 2010s.

Here is a list of common features controllable with HTTP:

Caching: How documents are cached can be controlled by HTTP. The server can instruct proxies and clients about what 
to cache and for how long. The client can instruct intermediate cache proxies to ignore the stored document.
Relaxing the origin constraint: To prevent snooping and other privacy invasions, Web browsers enforce strict 
separation between Web sites. Only pages from the same origin can access all the information of a Web page. 
Though such a constraint is a burden to the server, HTTP headers can relax this strict separation on the server side, 
allowing a document to become a patchwork of information sourced from different domains; there could even be 
security-related reasons to do so.

Authentication: Some pages may be protected so that only specific users can access them. Basic authentication may 
be provided by HTTP, either using the WWW-Authenticate and similar headers, or by setting a specific session using 
HTTP cookies.

Proxy and tunneling: Servers or clients are often located on intranets and hide their true IP address from other 
computers. HTTP requests then go through proxies to cross this network barrier. Not all proxies are HTTP proxies. 
The SOCKS protocol, for example, operates at a lower level. Other protocols, like ftp, can be handled by these proxies.
Sessions: Using HTTP cookies allows you to link requests with the state of the server. This creates sessions, 
despite basic HTTP being a state-less protocol. This is useful not only for e-commerce shopping baskets, but also 
for any site allowing user configuration of the output.

## HTTP flow

When a client wants to communicate with a server, either the final server or an intermediate proxy, 
it performs the following steps:

Open a TCP connection: The TCP connection is used to send a request, or several, and receive an answer. 
The client may open a new connection, reuse an existing connection, or open several TCP connections 
to the servers.

Send an HTTP message: HTTP messages (before HTTP/2) are human-readable. With HTTP/2, these simple messages 
are encapsulated in frames, making them impossible to read directly, but the principle remains the same. For example:

```
GET / HTTP/1.1
Host: developer.mozilla.org
Accept-Language: fr
```

Read the response sent by the server, such as:

```
HTTP/1.1 200 OK
Date: Sat, 09 Oct 2010 14:28:02 GMT
Server: Apache
Last-Modified: Tue, 01 Dec 2009 20:18:22 GMT
ETag: "51142bc1-7449-479b075b2891b"
Accept-Ranges: bytes
Content-Length: 29769
Content-Type: text/html

<!DOCTYPE html>… (here come the 29769 bytes of the requested web page)

```
Close or reuse the connection for further requests.
If HTTP pipelining is activated, several requests can be sent without waiting for the first response 
to be fully received. HTTP pipelining has proven difficult to implement in existing networks, 
where old pieces of software coexist with modern versions. HTTP pipelining has been superseded in HTTP/2 
with more robust multiplexing requests within a frame.

## HTTP Messages

METHOD
HEADERS



## Conclusion
HTTP is an extensible protocol that is easy to use. The client-server structure, combined with 
the ability to add headers, allows HTTP to advance along with the extended capabilities of the Web.

Though HTTP/2 adds some complexity by embedding HTTP messages in frames to improve performance, 
the basic structure of messages has stayed the same since HTTP/1.0. Session flow remains simple, 
allowing it to be investigated and debugged with a simple HTTP message monitor.





HTTP headers are pieces of information that accompany HTTP requests or responses, serving as metadata. They convey details about security and the body of the request or response. All header fields follow a standardized format, consisting of a case-insensitive string (the field name) followed by a colon and its corresponding value. The format of the value is dependent on the header field and its intended purpose.


###
The POST verb is most-often utilized to **create** new resources. On successful creation, return HTTP status 201, returning a Location header with a link to the newly-created resource with the 201 HTTP status.

POST is neither safe nor idempotent. It is therefore recommended for non-idempotent resource requests. Making two identical POST requests will most-likely result in two resources containing the same information.

Examples:

POST http://www.example.com/customers
POST http://www.example.com/customers/12345/orders

The HTTP GET method is used to read (or retrieve) a representation of a resource. In the “happy” (or non-error) path, GET returns a representation in XML or JSON and an HTTP response code of 200 (OK). In an error case, it most often returns a 404 (NOT FOUND) or 400 (BAD REQUEST).

According to the design of the HTTP specification, GET (along with HEAD) requests are used only to read data and not change it. Therefore, when used this way, they are considered safe. That is, they can be called without risk of data modification or corruption—calling it once has the same effect as calling it 10 times, or none at all. Additionally, GET (and HEAD) is idempotent, which means that making multiple identical requests ends up having the same result as a single request.

Do not expose unsafe operations via GET—it should never modify any resources on the server.

Examples:

GET http://www.example.com/customers/12345
GET http://www.example.com/customers/12345/orders
GET http://www.example.com/buckets/sample



PUT is most-often utilized for update capabilities, PUT-ing to a known resource URI with the request body containing the newly-updated representation of the original resource.

However, PUT can also be used to create a resource in the case where the resource ID is chosen by the client instead of by the server. In other words, if the PUT is to a URI that contains the value of a non-existent resource ID. Again, the request body contains a resource representation. Many feel this is convoluted and confusing. Consequently, this method of creation should be used sparingly, if at all.

Alternatively, use POST to create new resources and provide the client-defined ID in the body representation—presumably to a URI that doesn't include the ID of the resource (see POST below).

On successful update, return 200 (or 204 if not returning any content in the body) from a PUT. If using PUT for create, return HTTP status 201 on successful creation. A body in the response is optional—providing one consumes more bandwidth. It is not necessary to return a link via a Location header in the creation case since the client already set the resource ID.

PUT is not a safe operation, in that it modifies (or creates) state on the server, but it is idempotent. In other words, if you create or update a resource using PUT and then make that same call again, the resource is still there and still has the same state as it did with the first call.

If, for instance, calling PUT on a resource increments a counter within the resource, the call is no longer idempotent. Sometimes that happens and it may be enough to document that the call is not idempotent. However, it's recommended to keep PUT requests idempotent. It is strongly recommended to use POST for non-idempotent requests.

Examples:

PUT http://www.example.com/customers/12345
PUT http://www.example.com/customers/12345/orders/98765
PUT http://www.example.com/buckets/secret_stuff



DELETE
DELETE is used to delete a resource identified by a URI.

On successful deletion, might return HTTP status 200 (OK) along with a response body, or HTTP status 204 (NO CONTENT) with no response body. 



Examples:

DELETE http://www.example.com/customers/12345
DELETE http://www.example.com/customers/12345/orders
DELETE http://www.example.com/bucket/sample

[What is the Internet](https://www.youtube.com/watch?v=Dxcc6ycZ73M)
<BR/>
[The Internet: IP Addresses & DNS](https://www.youtube.com/watch?v=5o8CwafCxnU)
<BR/>
[The Internet: Wires, Cables & Wifi](https://www.youtube.com/watch?v=ZhEf7e4kopM)
<BR/>
[The Internet: Packets, Routing & Reliability](https://www.youtube.com/watch?v=AYdF7b3nMto)
<BR/>
[The Internet: HTTP & HTML](https://www.youtube.com/watch?v=kBXQZMmiA4s)
<BR/>
[Backend web development](https://www.youtube.com/watch?v=XBu54nfzxAQ)
<BR/>
[The Internet: Encryption & Public Keys](https://www.youtube.com/watch?v=ZghMPWGXexs&list=PLzdnOPI1iJNfMRZm5DDxco3UdsFegvuB7&index=7)

in the previous section, we learned that the Domain Name System (DNS) provides a human-friendly name for a Web server's IP address. This means that instead of having to remember a long string of numbers and letters, we can use a domain name, such as "www.example.com" to access a website.

However, even with the use of domain names, we still need a way to describe a specific resource within that server, such as a web page. This is where the Uniform Resource Locator, or URL, comes in.

A URL is a string of text that describes the location of a resource on the Internet. It typically includes the protocol used to access the resource, such as "http" or "https", the domain name or IP address of the server, and the location of the resource within the server, such as a file path or a specific page on a website.

For example, "https://www.example.com/about-us" is a URL that points to the "about-us" page on the "example.com" website.

Therefore, while DNS provides a way to easily remember and access a web server, URLs provide a way to identify and access specific resources within that server. It's important to remember that DNS and URLs work together to make it easy for us to access the resources on the Internet.




In previous section we learned Domain Name System (DNS) provides a human-friendly name for a Web server's IP address. Still we still need a way to describe a specific resource, such as a web page, within that server.

For this purpose, we use a Uniform Resource Locator, or URL for short. A URL is a string of text that describes the location of a resource on the Internet. It typically includes the protocol used to access the resource (such as "http" or "https"), the domain name or IP address of the server, and the location of the resource within the server (such as a file path or a specific page on a website).

For example, "https://www.example.com/about-us" is a URL that points to the "about-us" page on the "example.com" website.

Therefore, while DNS provides a way to easily remember and access a web server, URLs provide a way to identify and access specific resources within that server.




While IP addresses are enough to identify a host, but they are not very convenient for humans to remember.
To solve this, we the more userfriendly Domain Name System (DNS). DNS is a hierarcical  


While IP addresses are enough to identify devices on the Internet, they can be difficult for humans to remember.


Every device connected to the Internet has at least one IP address.

Just as your home has a unique address that is used to identify its location, devices connected to the Internet also have a unique address that is used to identify their location on the network.

This address is called an Internet Protocol Address, or IP Address for short. Every device that connects to the Internet, whether it's a computer, smartphone, or even a smart fridge, has its own unique IP address.
TCP/IP Illustrated, Volume 1, The Protocols


Devices used in private networks based on the TCP/IP protocols also require IP addresses. In either case, the forwarding procedures implemented by IP routers (see Chapter 5) use IP addresses to identify where traffic is going. IP addresses also indicate where traffic has come from. IP addresses are similar in some ways to telephone numbers, but whereas telephone numbers are often known and used directly by end users, IP addresses are often shielded from a user’s view by the Internet’s DNS (see Chapter 11), which allows most users to use names instead of numbers. 


Users are confronted with manipulating IP addresses when they are required to set up networks themselves or when the DNS has failed for some reason. To understand how the Internet identifies hosts and routers and delivers traffic between them, we must understand the role of IP addresses. We are therefore interested in their administration, structure, and uses.



A client-server architecture is a distributed computing model in which a client program sends requests to a server program over a network, and the server program responds to those requests. The client and the server are separate entities that run on different devices and communicate with each other using a standard protocol, such as HTTP or TCP/IP.  

As a backend developer, it is essential to have a solid understanding of design patterns and architecture. One of the fundamental concepts in web development is the client-server architecture. 

In this architecture, the client, such as a web browser, connects to a server and sends a request for data. The server then receives the request and responds by providing the requested information or data. 

Understanding this architecture is vital for designing and building efficient, scalable, and maintainable backend systems.


This architecture allows for the separation of concerns and enables the client and server to operate independently. Understanding this architecture is vital for designing and building efficient, scalable, and maintainable backend systems. Additionally, learning and understanding design patterns can help a developer to make more informed design decisions and write more organized, reusable code.

explain this as a teacher "As a backend developer, it's important to learn and understand design patterns and architecture.
The Web follows a client-server architecture, where the client connects and sends a request to the server to reuest data and the server receives the request and responds with requested data or information."


understand 


how the client and server interact and communicate with each other. The Web uses a client-server architecture, where one or more clients connect to the server to access resources and services.




applications that operate over a network are designed using a small number of specific patterns. Two of the most commonly used patterns are the client-server and peer-to-peer models. The client-server model is a centralized architecture, where one or more clients connect to a central server to access resources and services. 


As a backend developer, it is important to have a good understanding of the client-server architecture.

The client-server architecture is a way of organizing and structuring a software application, where the client and server interact with each other to complete a task. The client, in this case, refers to the device or application that makes a request to the server, and the server refers to the device or application that receives the request and responds with the requested data or information.

For example, when a user visits a website, their web browser (the client) sends a request to the server where the website is hosted. The server then responds by sending back the HTML, CSS, and JavaScript files that the browser uses to display the website to the user.

As a backend developer, it's important to understand how the client and server interact and communicate with each other. This includes understanding the different types of requests and responses, such as GET and POST requests, and how to handle and process them on the server side. Additionally, it's important to understand the different protocols and technologies that are used for communication between the client and server, such as HTTP and HTTPS.

By understanding the client-server architecture, backend developers can develop and maintain server-side code that is able to effectively handle client requests and respond with the appropriate data or information.


[](https://careerfoundry.com/en/blog/web-development/backend-developer-guide/)
# How the Web Works

but they can be difficult for humans to remember or use effectively. To address this issue, we use the Domain Name System (DNS) which allows for more user-friendly domain names to be associated with specific IP addresses.

"HTTP, or Hypertext Transfer Protocol, is a communication protocol used for transmitting data over the World Wide Web. It is the foundation of data communication for the web and is based on a request-response model.

When a client, such as a web browser, wants to request information from a server, it sends an HTTP request message. This message contains a method, such as GET or POST, that specifies the action to be taken, a URI or URL identifying the resource to be acted upon, and various headers providing additional information about the request.

The server then processes the request and sends an HTTP response message back to the client. This message contains a status code indicating the outcome of the request, such as 200 for a successful request or 404 for a page not found, as well as headers and a message body containing the requested information.

It's worth noting that HTTPS is the secure version

Today, we will begin exploring the fundamental concepts of the World Wide Web by examining the layered stack from the bottom up. Our focus will be on the network, or connectivity, layer - the backbone of the global network that enables computers to communicate with one another using numerical IP addresses, such as 172.217.13.132. Let's dive in and gain a deeper understanding of how this critical component functions.


IP addresses are enough to identify a host, but they are not very convenient for humans to remember or manipulate

applications that operate over a network are designed using a small number of specific patterns. Two of the most commonly used patterns are the client-server and peer-to-peer models. The client-server model is a centralized architecture, where one or more clients connect to a central server to access resources and services. 



The World Wide Web (WWW or Web) is a system of interlinked hypertext documents that are accessed via the Internet. It was created by Sir Tim Berners-Lee in 1989 while he was working at CERN, the European physics research organization. The Web uses a client-server model, where users interact with web browsers (clients) to request pages from web servers, which then return the requested information.

The foundation of the Web is the Hypertext Transfer Protocol (HTTP), which is used to transfer data between the client and server. Web pages are written in Hypertext Markup Language (HTML), which provides the structure and layout of the page. Cascading Style Sheets (CSS) are used to control the visual presentation of the page, and JavaScript is used to add interactive elements and dynamic behavior.

Web pages are identified by unique addresses called URLs (Uniform Resource Locators). When a user enters a URL into their web browser, the browser sends a request to the server for the corresponding page. The server then returns the HTML, CSS, and JavaScript files for the page, which the browser uses to render the page for the user.

The Web also relies on the use of unique identifiers called IP addresses to identify and locate computers on the Internet. The Domain Name System (DNS) is used to match domain names (such as www.example.com) to the corresponding IP addresses.

References:

Berners-Lee, T., Fielding, R., & Masinter, L. (2005). "Uniform Resource Identifiers (URI): Generic Syntax". Internet Engineering Task Force (IETF).
Deitel, P.J., Deitel, H.M. (2018) "Internet & World Wide Web: How to Program" Pearson.
W3C (2022) "Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content"
W3C (2022) "Cascading Style Sheets (CSS)
ECMA International (2022) "ECMAScript Language Specification"
IETF (2022) "Domain Name System (DNS) Parameters"


# Take 2
The web, or World Wide Web, is a vast network of interconnected documents and other resources, linked by hyperlinks and URLs. The web is built on top of the internet, which is a global network of connected computers.

When a user wants to access a webpage, they use a web browser, such as Chrome or Firefox, to send a request to the server where the webpage is located. The request is sent using the Hypertext Transfer Protocol (HTTP), which is the standard protocol for communication on the web.

The request contains the URL of the requested resource, as well as other information such as the browser's capabilities and the user's preferences. The server then processes the request and sends a response back to the client. The response typically includes the requested webpage, as well as any additional resources that are needed to display the page, such as images, stylesheets, and JavaScript files.

The web page itself is typically written in HTML, which stands for Hypertext Markup Language. HTML provides the structure and layout for the webpage, using a set of predefined tags and attributes. CSS, or Cascading Style Sheets, is used to define the visual presentation of the webpage, such as colors, fonts, and layouts. JavaScript is a programming language that can be used to make web pages interactive, and can dynamically update the page's content and layout.

When the browser receives the response, it parses the HTML, CSS and JavaScript, then it applies the style and layout rules to the HTML and Runs the JavaScript. This is how the webpage is rendered and displayed to the user.

The web is a constantly evolving ecosystem, and new technologies and standards are constantly being developed to enhance the functionality and capabilities of the web. The web is accessible from anywhere in the world as long as there is an internet connection.

In short, the web is a vast network of interconnected documents and resources that can be accessed by a user via a web browser. The web pages are typically written in HTML, CSS, and JavaScript, and the browser uses these technologies to render the webpage and display it to the user. All this process happens through the interaction between the browser and the web server, which follow the HTTP protocol to send requests and responses, and all this happening over the internet.

[Internet](https://youtu.be/Dxcc6ycZ73M?list=PLzdnOPI1iJNfMRZm5DDxco3UdsFegvuB7)

[History of the World Wide Web](https://www.hanselminutes.com/292/history-of-http-and-the-world-wide-web-with-henrik-frystyk-nielsen)

[How the web works](https://www.preethikasireddy.com/post/how-the-web-works-a-primer-for-newcomers-to-web-development-or-anyone-really)<BR/>
[Structure of a web aplication](https://www.preethikasireddy.com/post/how-the-web-works-part-ii-client-server-model-the-structure-of-a-web-application)

# Backend Developer
A backend developer is a software developer who specializes in the development of the server-side of web applications. They are responsible for building the infrastructure and logic that supports the front-end user interface, and they work closely with front-end developers and other members of the development team to create a cohesive and functional web application.


The server-side logic is the code that runs on the server and performs tasks such as processing user requests, querying the database for data, and sending responses back to the client.


Backend developers use programming languages such as Java, Python, Ruby, or Node.js to write this logic. They also use frameworks and technologies that are specific to the language and platform they are working on, such as Express.js for Node.js or Rails for Ruby.

The database is the storage system where the web application stores and retrieves data. Backend developers work closely with database administrators to design and optimize the database schema for the web application. They also write code to interact with the database using languages such as SQL and NoSQL.

APIs are the bridge between the front-end and the back-end of the web application. They allow the front-end to send requests to the server and receive responses in a structured format. Backend developers are responsible for designing and implementing the APIs, using technologies such as REST or GraphQL.

In summary, Backend development refers to the server-side software that is responsible for how a website works, in other words, how the database communicates with the browser through a written code. It includes server-side logic, database design and management, and API design and implementation. Backend developers work closely with front-end developers and database administrators to ensure that the web application is able to function correctly.


A backend developer's responsibilities typically include:

Designing and implementing server-side logic and APIs that handle client requests and handle data storage, retrieval, and manipulation.
Writing code that is optimized for performance and scalability.
Implementing security measures to protect the application and its data from unauthorized access and malicious attacks.
Integrating with third-party systems and services as needed, such as databases, storage systems, and external APIs.
Testing and debugging code to ensure it is free of errors and runs smoothly.
Backend developers are often proficient in server-side languages such as Node.js, Python, Ruby, Java, Go and others. They often work with databases such as MySQL, PostgreSQL, MongoDB, and others, and they may also use web framework such as Ruby on Rails, Django, Express.js, Flask, etc.

Some of the core skills that a backend developer should possess include a solid understanding of web development concepts, an understanding of database design and management, experience with server-side languages, and knowledge of security and performance best practices. Additionally, backend developers should have good problem-solving skills, the ability to work well in a team environment, and strong communication skills.

Overall, the role of a backend developer is a crucial one in the web development process as they build the foundation of a web application that supports the frontend interfaces, they are responsible for handling data, implementing security measures and ensuring a good performance and scalability.

# Client Server
A client-server architecture is a distributed computing model in which a client program sends requests to a server program over a network, and the server program responds to those requests. The client and the server are separate entities that run on different devices, and they communicate with each other using a standard protocol, such as HTTP or TCP/IP.

The client is typically a program or device that runs on a user's device, such as a web browser, a mobile app, or a desktop application. The client is responsible for presenting the user interface, handling user input, and displaying the data that is returned from the server.

The server, on the other hand, is a program or device that runs on a separate machine and is responsible for handling the client's requests and providing the necessary data or services. The server can be a powerful computer or a cluster of computers that are dedicated to providing a specific service, such as storing and retrieving files or managing a database.

The server and client communicate through network protocols, which is a set of standard rules that dictate how data should be exchanged over the network. These protocols include HTTP and HTTPS for the web, FTP for file transfer, SMTP for sending emails, and DNS for domain name resolution.

In short, client-server architecture is a distributed computing model in which a client program (running on the client device) sends requests to the server program (running on the server device) over the network, the server then processes the request and sends a response back to the client. This communication between client and server happens by following a standard set of protocol rules, they are specific to the type of service.

# WEB
A web client-server architecture is a specific implementation of the client-server architecture that is used to deliver web pages and other content over the internet. It involves a web client, such as a web browser, and a web server, which is a program that runs on a computer or cluster of computers that are connected to the internet.

The web client, such as a web browser, runs on the user's device and is responsible for requesting web pages and other content from the server. When a user types a URL or clicks on a link, the web browser sends a request to the server using the HTTP protocol. The request includes information such as the URL of the requested resource and any additional query parameters that may be needed.

The web server, on the other hand, is a program that runs on a separate machine and is responsible for handling the client's requests and providing the necessary data or services. The server listens for incoming requests, processes them, and sends a response back to the client. The response typically includes the requested web page, along with any additional data or resources that are needed to display the page.

The server can also be connected to a database and can access the data from it, then it sends that data to the client.

In short, the web client-server architecture is a way in which web pages and other content are delivered over the internet. The web client (typically a web browser) runs on the user's device, and requests web pages from the web server which runs on a separate machine and processes the request, then sends the response back to the client. The communication between the client and server happens using the HTTP protocol which defines a set of standard rules for data to be exchanged over the network.

# DNS
The Domain Name System (DNS) is a hierarchical and decentralized naming system for computers, services, or any resource connected to the Internet or a private network. It translates more easily memorized domain names, such as www.example.com, to the numerical IP addresses, such as 192.0.2.1, that computers use to identify each other on the network.

DNS is an essential component of the Internet infrastructure, as it allows users to access websites and other resources using human-readable domain names rather than IP addresses. Without DNS, users would have to remember and manually enter the IP addresses of the websites they wish to visit, which would be a cumbersome and error-prone process.

The DNS system is organized into a hierarchy of domain names, with the top-level domains (TLDs) being the highest level. TLDs such as .com, .org, and .edu are managed by the Internet Corporation for Assigned Names and Numbers (ICANN). Below the TLDs, there can be multiple levels of subdomains, such as www.example.com, where "example" is the second-level domain and "www" is the third-level domain.

DNS works by using a distributed database of domain name-to-IP address mappings, called the DNS namespace. This database is distributed across multiple DNS servers, and each server is responsible for a specific portion of the namespace. When a client computer requests a website using its domain name, the request is first sent to a local DNS resolver, which then queries the appropriate DNS server for the IP address of the requested website. This process is called DNS resolution.

It's important to note that the DNS system is not only used for websites but also for other types of network resources such as email servers, FTP servers, and more. In addition, DNS can be used for load balancing and for providing failover capabilities for servers.

References:

IETF (2022) "Domain Name System (DNS) Parameters"
Albitz, P. and Liu, C. (2012) "DNS and BIND" O'Reilly Media
Liu, Cricket (2017) "DNS and BIND" O'Reilly Media
Roessler, P. (2018) "DNS and BIND on IPv6" O'Reilly Media





A domain name is a unique string of characters that identifies a website or a specific webpage on the internet. Think of it as an address for a website, just like how a physical address is used to identify a house or building.

For example, the domain name for Google is "google.com." In this case, "google" is the specific name chosen by the company, and ".com" is the top-level domain (TLD), which indicates the type of organization or purpose of the website. Other examples of TLDs include ".org" for organizations, ".edu" for educational institutions, and ".gov" for government agencies.

When a user types a domain name into their web browser, their computer sends a request to a special server called a Domain Name System (DNS) server, which translates the domain name into an IP address that the computer can use to locate the website on the internet. This process is similar to how a phone book works - it maps names to numbers.

In simple terms, a Domain name serves as human-readable address that point to a web page, the DNS server converts it to an IP address (that a computer can read) which points the browser to the right web server.

One can purchase domain names through a registrar which is accredited by the Internet Corporation for Assigned Names and Numbers (ICANN), a non-profit organization that oversees the allocation of domain names. Once you purchase a domain, you must renew it annually to maintain control over it.

Overall, domain names provide a convenient and memorable way for users to find and access websites on the internet.

The Domain Name System (DNS) is a hierarchical and decentralized naming system for computers, services, or any resource connected to the Internet or a private network. It translates more easily memorized domain names, such as www.example.com, to the numerical IP addresses, such as 192.0.2.1, that computers use to identify each other on the network.

DNS is an essential component of the Internet infrastructure, as it allows users to access websites and other resources using human-readable domain names rather than IP addresses. Without DNS, users would have to remember and manually enter the IP addresses of the websites they wish to visit, which would be a cumbersome and error-prone process.

The DNS system is organized into a hierarchy of domain names, with the top-level domains (TLDs) being the highest level. TLDs such as .com, .org, and .edu are managed by the Internet Corporation for Assigned Names and Numbers (ICANN). Below the TLDs, there can be multiple levels of subdomains, such as www.example.com, where "example" is the second-level domain and "www" is the third-level domain.

DNS works by using a distributed database of domain name-to-IP address mappings, called the DNS namespace. This database is distributed across multiple DNS servers, and each server is responsible for a specific portion of the namespace. When a client computer requests a website using its domain name, the request is first sent to a local DNS resolver, which then queries the appropriate DNS server for the IP address of the requested website. This process is called DNS resolution.

It's important to note that the DNS system is not only used for websites but also for other types of network resources such as email servers, FTP servers, and more. In addition, DNS can be used for load balancing and for providing failover capabilities for servers.

References:

IETF (2022) "Domain Name System (DNS) Parameters"
Albitz, P. and Liu, C. (2012) "DNS and BIND" O'Reilly Media
Liu, Cricket (2017) "DNS and BIND" O'Reilly Media
Roessler, P. (2018) "DNS and BIND on IPv6" O'Reilly Media










[IBM What is DNS](https://www.youtube.com/watch?v=nyH0nYhMW9M)<BR/>
[What is DNS]([https://aws.amazon.com/route53/what-is-dns/](https://www.youtube.com/watch?v=e2xLV7pCOLI))

# URL
A URL, or uniform resource locator, is the specific address of a webpage or resource on the internet. It is used to identify and locate a specific file or webpage on the World Wide Web. A URL typically consists of several parts, including the protocol, the domain name, and the file path.

For example, the URL for the Google homepage is "https://www.google.com". In this example, "https" is the protocol, which indicates that the webpage is being accessed securely. "google.com" is the domain name, and the file path, "/" implies it's the home page.

The protocol specifies the way a browser should retrieve the information from the server. For instance, "http" and "https" are the two most common protocols used on the internet, "http" stands for HyperText Transfer Protocol and "https" stands for HTTP Secure, which uses SSL/TLS certificate to encrypt data in transit to provide more secure communication.

You can also include additional information in a URL by using query parameters, which are appended to the end of the file path. These query parameters are used to specify specific values for certain parameters on the webpage, such as specifying a search query or filtering results.

For instance, the URL of a search query in Google "https://www.google.com/search?q=example" , Here "q" is the parameter and "example" is the value.

In short, a URL is a specific address that is used to locate and retrieve a resource on the internet. It includes the protocol, domain name, file path, and sometimes query parameters that specify additional information about the resource being requested.


# IP addresses and Port Numbers
An IP address is a numerical label assigned to each device connected to a computer network that uses the Internet Protocol for communication. It serves two main functions: identifying the host or network interface, and providing the location of the host in the network.

IP addresses come in two types, IPv4 and IPv6. IPv4 addresses are 32-bit numbers, written in the form of four 8-bit numbers separated by periods (e.g., 192.168.1.1), and can support up to 4.3 billion unique addresses. IPv6 addresses, on the other hand, are 128-bit numbers, written in the form of eight 16-bit numbers separated by colons (e.g., 2001:0db8:85a3:0000:0000:8a2e:0370:7334) and can support up to 340 undecillion unique addresses.

In addition to the IP address, computers on a network use port numbers to identify different applications and services running on a device. A port number is a 16-bit unsigned integer that identifies a specific process or service running on a host. The combination of an IP address and a port number is called a socket, which is used to uniquely identify a specific service or process on a network.

For example, when you type "https://www.google.com" in your browser, the browser sends a request to the server with the IP address of the server and port number 443 (default for HTTPS), the server then processes the request and sends back the response to the browser via the IP address and port number provided in the request.

Common port numbers that are used by well-known services include 80 for HTTP, 443 for HTTPS, 22 for SSH, and 21 for FTP. However, any port number above 1024 can be used for custom or non-registered services.

In short, an IP address is a unique numerical label that is assigned to each device connected to a network, it's used to identify the host and provide its location in the network. A port number is a 16-bit unsigned integer that is used to identify a specific process or service running on a device, it's combined with the IP address to create a socket, that identifies a unique service or process on a network.

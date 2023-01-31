# Dynamic Host Configuration Protocol

## DHCP (Client)
In a DHCP (Dynamic Host Configuration Protocol) communication, the client sends a broadcast message called a "DHCP Discover"
to locate a DHCP server on the network. The "DHCP Discover" message includes information about the client, such as its host 
name and MAC address, as well as any DHCP options that the client is capable of using.

The client does not send any specific option for DNS configuration as the DHCP server is responsible for sending the 
required information to the client. The client simply requests that the DHCP server provide it with the necessary 
configuration information. The DHCP server then responds with a "DHCP Offer" message, which includes the assigned 
IP address and other network configuration information, including the DNS server information specified in DHCP options 
such as Option 6 (DNS Servers), Option 15 (Domain Name), or Option 81 (FQDN). The client then acknowledges the offer
and completes the DHCP process by sending a "DHCP Request" message, after which the DHCP server sends a "DHCP Ack" 
message to confirm the client's configuration.


## DHCP (Server)
DHCP (Dynamic Host Configuration Protocol) options are fields in DHCP packets that convey specific information to clients. 
The DHCP options used for configuring DNS servers include:

Option 6 (DNS Servers): This option specifies one or more IP addresses of DNS servers to be used by the client. 
The IP addresses are listed in order of priority.

- Option 15 (Domain Name): This option specifies the domain name to be used by the client for host name-to-IP address 
resolution.
- Option 81 (FQDN, Fully Qualified Domain Name): This option is used to provide the client with its fully-qualified domain name and its corresponding domain name servers.
Reference:

- IETF RFC 2132: "DHCP Options and BOOTP Vendor Extensions". This RFC specifies the format of DHCP options and the options available for use by DHCP clients and servers.
- IETF RFC 4702: "The Use of DHCP Option 43 and Option 60 for Lightweight Access Point Discovery and Configuration". This RFC specifies the use of DHCP options 43 and 60 for configuring lightweight access points.
It is important to note that DHCP options are not standardized across all DHCP implementations, and the availability and use of specific options may vary. However, options 6, 15, and 81 are commonly used for configuring DNS servers in DHCP.

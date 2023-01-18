# dnsmasq

```
dnsmasq --address=/skyramp.test/10.0.1.71 \
  --no-daemon \
  --listen-address=0.0.0.0 \
  --interface=eth0 \
  --log-queries \
  --no-resolv \
  --no-hosts \
  --server=8.8.8.8 \
  --server=8.8.4.4

```

```
dnsmasq --address=/skyramp.test/10.0.1.71 \
  --no-daemon \
  --listen-address=0.0.0.0 \
  --interface=eth0 \
  --log-queries \
  --no-hosts 
```

```
docker run -it pellepedro/dnsmasq \
  --no-resolve \
  --port=8053 --keep-in-foreground --no-daemon --no-hosts --log-queries \
  --address=/skyramp.test/203.0.113.2 \
```
```
docker run -it pellepedro/dnsmasq \
  --port=8053 --keep-in-foreground --no-daemon --no-hosts --log-queries \
  --address=/skyramp.test/203.0.113.2 \
  --address=/example.com/203.0.113.129 \
  --host-record=example.com,198.51.100.111 \
  --host-record=host.example.com,198.51.100.222 \
  --cname=alias.example.com,host.example.com \
  --cname=alias2.example.com,www.google.com \

```
## Netplan
```
network:
  version: 2
  renderer: networkd
  ethernets:
    eth0:
      dhcp4: no
      nameservers:
          addresses: [172.17.0.2]
```

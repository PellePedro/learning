# dnsmasq

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

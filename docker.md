[Private Registry](https://iceburn.medium.com/docker-private-registry-lets-encrypt-on-ubuntu-18-04-b310f79d116e)

```
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
    
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
  "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) \
  stable"

sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io
sudo usermod -aG docker ${USER}
```

```
sudo add-apt-repository ppa:certbot/certbot -y
sudo apt update
sudo apt install certbot -y
```

```
sudo su
certbot certonly --standalone --preferred-challenges http --non-interactive  --staple-ocsp --agree-tos -m admin@testdomain.com -d registry.cequoi.ca
```
Configure renewal of certificate
```
cat <<EOF > /etc/cron.d/letencrypt
SHELL=/bin/sh
PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
30 2 * * 1 root /usr/bin/certbot renew >> /var/log/letsencrypt-renew.log && cd /etc/letsencrypt/live/registry.cequoi.ca && cp privkey.pem domain.key && cat cert.pem chain.pem > domain.crt && chmod 777 domain.*
EOF
cd /etc/letsencrypt/live/registry.cequoi.ca && \
cp privkey.pem domain.key && \
cat cert.pem chain.pem > domain.crt && \
chmod 777 domain.*
mkdir -p /mnt/docker-registry
```
Run docker
```
docker run --entrypoint htpasswd registry:latest -Bbn pellepedro pellepedro2022 > /mnt/docker-registry/passfile
docker run -d -p 443:5000 --restart=always --name registry \
  -v /etc/letsencrypt/live/registry.cequoi.ca:/certs \
  -v /mnt/docker-registry:/var/lib/registry \
  -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \
  -e REGISTRY_HTTP_TLS_KEY=/certs/domain.key \
  -e REGISTRY_AUTH=htpasswd \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e REGISTRY_AUTH_HTPASSWD_PATH=/var/lib/registry/passfile \
  registry:2.7.1
```

```
docker login -u pellepedro -p pellepedro2022 registry.ceqoia.ca
```


[Building multi-platform Images](https://github.com/docker/buildx#building-multi-platform-images)


[buildx bake](https://github.com/docker/buildx#buildx-bake-options-target)

```sh
docker buildx bake --push --set *.platform=linux/amd64,linux/arm64`


docker run -it -p 53:53/udp public.ecr.aws/j1n2c2p2/rampup/dnsmasq:v0.4.0 dnsmasq --no-resolv --keep-in-foreground --no-hosts --listen-address=0.0.0.0 --bind-interfaces --address=/skyramp.test/127.0.0.1


```

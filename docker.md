[Private Registry](https://iceburn.medium.com/docker-private-registry-lets-encrypt-on-ubuntu-18-04-b310f79d116e)


[Building multi-platform Images](https://github.com/docker/buildx#building-multi-platform-images)


[buildx bake](https://github.com/docker/buildx#buildx-bake-options-target)

```sh
docker buildx bake --push --set *.platform=linux/amd64,linux/arm64`


docker run -it -p 53:53/udp public.ecr.aws/j1n2c2p2/rampup/dnsmasq:v0.4.0 dnsmasq --no-resolv --keep-in-foreground --no-hosts --listen-address=0.0.0.0 --bind-interfaces --address=/skyramp.test/127.0.0.1


```

[Private Registry](https://iceburn.medium.com/docker-private-registry-lets-encrypt-on-ubuntu-18-04-b310f79d116e)


[Building multi-platform Images](https://github.com/docker/buildx#building-multi-platform-images)


[buildx bake](https://github.com/docker/buildx#buildx-bake-options-target)

```sh
docker buildx bake --push --set *.platform=linux/amd64,linux/arm64`
```

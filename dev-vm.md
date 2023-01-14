
## Install Skyramp dependencies
```
sudo apt-get update -yq
sudo apt-get install -y --no-install-recommends \
btrfs-progs \
crun \
git \
make \
go-md2man \
iptables \
libassuan-dev \
libbtrfs-dev \
libc6-dev \
libdevmapper-dev \
libglib2.0-dev \
libgpgme-dev \
libgpg-error-dev \
libprotobuf-dev \
libprotobuf-c-dev \
libseccomp-dev \
libselinux1-dev \
libsystemd-dev \
pkg-config \
uidmap
       
```

## Install Golang
```
GO_VERSION="1.19.4"
GOLANG_TAR=linux-amd64.tar.gz
GOLANG_DOWNLOAD_URL=https://dl.google.com/go/go${GO_VERSION}.${GOLANG_TAR}
curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz
sudo tar -C /usr/local -xzf golang.tar.gz
rm golang.tar.gz
sudo ln -sf /usr/local/go/bin/* /usr/local/bin
rm -rf ${GOLANG_TAR}

```

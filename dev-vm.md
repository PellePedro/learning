
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

## Core packages
```
sudo apt-get install \
  bash            \
  bind-tools      \
  build-base      \
  ca-certificates \
  curl            \
  fzf             \
  git             \
  libstdc++       \
  net-tools       \
  nnn             \
  ripgrep         \
  vim             \
  zsh             \
  python3         \
  python3-dev

ln -sf python3 /usr/bin/python \
    && python -m ensurepip \
    && pip3 install --no-cache --upgrade pip setuptools codespell 
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


## Install go tools
```
go install github.com/jesseduffield/lazygit@latest
go install github.com/fatih/gomodifytags@latest
go install github.com/josharian/impl@latest
go get -u github.com/cweill/gotests/...
go get -u github.com/koron/iferr

go install honnef.co/go/tools/cmd/staticcheck@2022.1

# Node
curl -fsSL https://deb.nodesource.com/setup_19.x | sudo -E bash - &&\
sudo apt-get install -y nodejs
```



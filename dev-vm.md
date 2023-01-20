
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
  ca-certificates \
  curl            \
  fzf             \
  git             \
  nnn             \
  ripgrep         \
  vim             \
  zsh             \
  python3         \
  python3-dev     \
  python3-pip

sudo ln -sf python3 /usr/bin/python \
    && python -m ensurepip \
    && pip3 install --no-cache --upgrade pip setuptools codespell 
```

## Install Docker
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

## Install docker compose
[ref](https://github.com/josepmariatuset/jenkins-sessions/blob/c1d3d9031bef47e1851edac5333c5fb28a5d84e8/Dockerfile)
```
curl --fail -sL https://api.github.com/repos/docker/compose/releases/latest| grep tag_name | cut -d '"' -f 4 | tee /tmp/compose-version
sudo mkdir -p /usr/lib/docker/cli-plugins 
sudo  curl --fail -sL -o /usr/lib/docker/cli-plugins/docker-compose https://github.com/docker/compose/releases/download/$(cat /tmp/compose-version)/docker-compose-$(uname -s)-$(uname -m)
sudo chmod +x /usr/lib/docker/cli-plugins/docker-compose
sudo ln -s /usr/lib/docker/cli-plugins/docker-compose /usr/bin/docker-compose
rm /tmp/compose-version
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

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
go install honnef.co/go/tools/cmd/staticcheck@2022.1
go install github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/tools/gopls@latest

# Node
curl -fsSL https://deb.nodesource.com/setup_19.x | sudo -E bash - &&\
sudo apt-get install -y nodejs
```





## Install neovim
```
cat <<EOF > Dockerfile.nvim
FROM debian:bullseye as builder

#ARG NVIM_RELEASE=release
ARG NVIM_RELEASE=master

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && apt-get install --no-install-recommends -y \
    apt-transport-https \
    autoconf \
    automake \
    clang \
    cmake \
    curl \
    doxygen \
    g++ \
    gettext \
    git \
    gperf \
    libluajit-5.1-dev \
    libmsgpack-dev \
    libtermkey-dev \
    libtool \
    libtool-bin \
    libunibilium-dev \
    libutf8proc-dev \
    libuv1-dev \
    libvterm-dev \
    luajit \
    luarocks \
    make \
    ninja-build \
    pkg-config \
    unzip \
    ca-certificates

RUN luarocks build mpack && \
    luarocks build lpeg      && \
    luarocks build inspect

# Build neovim from source
ENV CMAKE_EXTRA_FLAGS="-DENABLE_JEMALLOC=OFF" \
  CMAKE_BUILD_TYPE="RelWithDebInfo"

RUN git clone https://github.com/neovim/neovim.git --branch \$NVIM_RELEASE \
  && cd neovim \
  && make \
  && make install

FROM scratch as artifact
COPY --from=builder /usr/local/bin/ /usr/local/bin
COPY --from=builder /usr/local/share/nvim/ /usr/local/share/nvim/
EOF
```

## Build and install nvim
```
sudo DOCKER_BUILDKIT=1 docker build --target=artifact --output type=local,dest=/ -f Dockerfile.nvim  .
```

## Install lunarvim
```
bash <(curl -s https://raw.githubusercontent.com/lunarvim/lunarvim/master/utils/installer/install.sh)
```

```
alias lvim=/home/ubuntu/.local/bin/lvim
```


## Install AWS CLI
```
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

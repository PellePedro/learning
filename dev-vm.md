
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

go install honnef.co/go/tools/cmd/staticcheck@2023.1.2
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
go install honnef.co/go/tools/cmd/staticcheck@2022.1
go install github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

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
sudo pip install aws-mfa
rm awscliv2.zip
```

## Login to AWS registry
```
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/j1n2c2p2
```

## Install kubectl
```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin
```

## Install skaffold
```
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/ && \
rm skaffold
```

## Install kind
```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.17.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```
## Install Helm
```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod +x get_helm.sh
sudo ./get_helm.sh
helm version
```

## Install k9s
```
# NOTE: The dev version will be in effect!
go install github.com/derailed/k9s@latest
```

## Install delta
```
curl -Lo ./git-delta_0.15.1_amd64.deb https://github.com/dandavison/delta/releases/download/0.15.1/git-delta_0.15.1_amd64.deb
sudo dpkg -i ./git-delta_0.15.1_amd64.deb 
rm ./git-delta_0.15.1_amd64.deb 
```

```
[user]
	name = pellepedro
	email = Per.Pettersson@gmail.com
[filter "lfs"]
	smudge = git-lfs smudge -- %f
	process = git-lfs filter-process
	required = true
	clean = git-lfs clean -- %f
[init]
	defaultBranch = main

[delta]
    features = side-by-side line-numbers decorations
    syntax-theme = Dracula
    navigate = true    # use n and N to move between diff sections
    light = false      # set to true if you're in a terminal w/ a light background color (e.g. the default macOS terminal)
    plus-style = syntax "#003800"
    minus-style = syntax "#3f0001"

[delta "decorations"]
    commit-decoration-style = bold yellow box ul
    file-style = bold yellow ul
    file-decoration-style = none
    hunk-header-decoration-style = cyan box ul

[delta "line-numbers"]
    line-numbers-left-style = cyan
    line-numbers-right-style = cyan
    line-numbers-minus-style = 124
    line-numbers-plus-style = 28

[apager]
    diff = delta
    log = delta
    reflog = delta
    show = delta

```

```
git:
    paging:
    ▏ colorArg: always
    ▏ pager: delta --dark --paging=never
    pull:
    ▏ mode: 'rebase'
  gui:
    theme:
    ▏ activeBorderColor:
    ▏ ▏ - blue
    ▏ ▏ - bold
    ▏ inactiveBorderColor:
    ▏ ▏ - white
    ▏ optionsTextColor:
    ▏ ▏ - blue
    ▏ selectedLineBgColor:
    ▏ ▏ - default # set to `default` to have no background colour
    ▏ selectedRangeBgColor:
    ▏ ▏ - default
    ▏ cherryPickedCommitBgColor:
    ▏ ▏ - cyan
    ▏ cherryPickedCommitFgColor:
    ▏ ▏ - blue
    ▏ unstagedChangesColor:
    ▏ ▏ - red
    ▏ ▏
    showFileTree: true # for rendering changes files in a tree format
    showListFooter: false # for seeing the '5 of 20' message in list panels
    showRandomTip: false
    showBottomLine: false # for hiding the bottom information line (unless it has important information to tell you)
    showCommandLog: true
    showIcons: true

  disableStartupPopups: true
  notARepository: 'skip' # one of: 'prompt' | 'create' | 'skip'
  os:
    editCommand: 'lvim'
```








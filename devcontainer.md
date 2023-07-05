# Devcontainer

[rjfmachado kubectl](https://github.com/rjfmachado/devcontainer-features/actions)   



## cli
```
npm install -g @devcontainers/cli
```

## devcontainer
```
devcontainer build --workspace-folder sample-microservices
```


## Install Skyramp
```
#!/bin/sh
set -e
echo "Installing feature 'skyramp'"
CI=true bash -c "$(curl -fsSL  https://skyramp-public.s3.us-west-2.amazonaws.com/installer/install.sh)"
```

# Repos

[vankichi](https://github.com/vankichi/dotfiles/tree/master/dockers)

https://github.com/okteto/devenv

[pulumi-nutanix](https://github.com/s0nyguy/pulumi-nutanix/blob/main/.devcontainer/Dockerfile)


## Ubuntu

```
go install github.com/fatih/gomodifytags@latest
go install github.com/josharian/impl@latest
go get -u github.com/cweill/gotests/...
go get -u github.com/koron/iferr

go install honnef.co/go/tools/cmd/staticcheck@2022.1

# Node
curl -fsSL https://deb.nodesource.com/setup_19.x | sudo -E bash - &&\
sudo apt-get install -y nodejs

# Lunarvim
LV_BRANCH='release-1.2/neovim-0.8' bash <(curl -s https://raw.githubusercontent.com/lunarvim/lunarvim/master/utils/installer/install.sh)


# Python
```

### Install Go on Ubuntu
```bash
GO_VERSION="1.19.4"
GOLANG_TAR=linux-amd64.tar.gz
GOLANG_DOWNLOAD_URL=https://dl.google.com/go/go${GO_VERSION}.${GOLANG_TAR}
curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz
sudo tar -C /usr/local -xzf golang.tar.gz
rm golang.tar.gz
sudo ln -sf /usr/local/go/bin/* /usr/local/bin
rm -rf ${GOLANG_TAR}
```

### vscode server
```
# to simplify development within the RoSys container we preinstall vs code server and extensions
RUN curl -sSL https://gist.githubusercontent.com/b01/0a16b6645ab7921b0910603dfb85e4fb/raw/5186ea07a06eac28937fd914a9c8f9ce077a978e/download-vs-code-server.sh | sed "s/server-linux-x64/server-linux-$(dpkg --print-architecture)/" | sed "s/amd64/x64/" | bash
ENV VSCODE_SERVER=/root/.vscode-server/bin/*/server.sh
RUN $VSCODE_SERVER --install-extension ms-python.vscode-pylance \
    $VSCODE_SERVER --install-extension ms-python.python \
    $VSCODE_SERVER --install-extension himanoa.python-autopep8 \
    $VSCODE_SERVER --install-extension esbenp.prettier-vscode \
```    
```
export commit_id=b380da4ef1ee00e224a15c1d4d9793e27c2b6302
wget curl -sSL "https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable" -o vscode-server-linux-x64.tar.gz
mkdir -p ~/.vscode-server/bin/${commit_id}
tar zxvf vscode-server-linux-x64.tar.gz -C ~/.vscode-server/bin/${commit_id} --strip 1
touch ~/.vscode-server/bin/${commit_id}/0
```

```
https://update.code.visualstudio.com/commit:b380da4ef1ee00e224a15c1d4d9793e27c2b6302/

ENV commit_id=b380da4ef1ee00e224a15c1d4d9793e27c2b6302

# Download url is: https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable
curl -sSL "https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable" -o vscode-server-linux-x64.tar.gz

mkdir -p ~/.vscode-server/bin/${commit_id}
tar zxvf vscode-server-linux-x64.tar.gz -C ~/.vscode-server/bin/${commit_id} --strip 1
touch ~/.vscode-server/bin/${commit_id}/0


```


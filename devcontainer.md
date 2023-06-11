# Devcontainers

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
export commit_id=b380da4ef1ee00e224a15c1d4d9793e27c2b6302
wget curl -sSL "https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable" -o vscode-server-linux-x64.tar.gz
mkdir -p ~/.vscode-server/bin/${commit_id}
tar zxvf vscode-server-linux-x64.tar.gz -C ~/.vscode-server/bin/${commit_id} --strip 1
touch ~/.vscode-server/bin/${commit_id}/0
```

```
ENV commit_id=b380da4ef1ee00e224a15c1d4d9793e27c2b6302

# Download url is: https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable
curl -sSL "https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable" -o vscode-server-linux-x64.tar.gz

mkdir -p ~/.vscode-server/bin/${commit_id}
tar zxvf vscode-server-linux-x64.tar.gz -C ~/.vscode-server/bin/${commit_id} --strip 1
touch ~/.vscode-server/bin/${commit_id}/0


```


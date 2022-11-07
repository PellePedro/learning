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

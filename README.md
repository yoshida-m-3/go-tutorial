# go-tutorial

Go学習用

A Tour of Goの写経

## vim setting

### install go

```
brew install go
```

### install formatter

install [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

```
go install golang.org/x/tools/cmd/goimports@latest
```

#### setup environment

.zshrc

```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

# 重複パスを登録しない
typeset -U path PATH
```
